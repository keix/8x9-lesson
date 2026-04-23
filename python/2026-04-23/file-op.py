from typing import List, Dict, Optional
import os


class FileRecord:
    """Represents a file entry with name and size."""

    def __init__(self, name: str, size: int) -> None:
        self.name = name
        self.size = size

    def __repr__(self) -> str:
        return f"{self.name} ({self.size} bytes)"


def list_files(directory: str) -> List[FileRecord]:
    """List files in a directory."""
    records: List[FileRecord] = []
    try:
        for entry in os.scandir(directory):
            if entry.is_file():
                size = entry.stat().st_size
                records.append(FileRecord(entry.name, size))
    except FileNotFoundError:
        print("Directory not found.")
    except PermissionError:
        print("Permission denied.")
    return records


def filter_files(records: List[FileRecord], keyword: str) -> List[FileRecord]:
    """Filter files by keyword."""
    return [r for r in records if keyword.lower() in r.name.lower()]


def sort_files(records: List[FileRecord], reverse: bool = False) -> List[FileRecord]:
    """Sort files by size."""
    return sorted(records, key=lambda r: r.size, reverse=reverse)


def summarize(records: List[FileRecord]) -> Dict[str, int]:
    """Summarize file statistics."""
    total_size = sum(r.size for r in records)
    return {"count": len(records), "total_size": total_size}


def display(records: List[FileRecord]) -> None:
    """Display file list."""
    for r in records:
        print(r)


def read_file(path: str) -> Optional[str]:
    """Read file content."""
    try:
        with open(path, "r", encoding="utf-8") as f:
            return f.read()
    except FileNotFoundError:
        print("File not found.")
    except UnicodeDecodeError:
        print("Cannot decode file.")
    return None


def write_file(path: str, content: str) -> bool:
    """Write content to file."""
    try:
        with open(path, "w", encoding="utf-8") as f:
            f.write(content)
        return True
    except IOError:
        print("Write failed.")
        return False


def append_log(log_path: str, message: str) -> None:
    """Append message to log file."""
    try:
        with open(log_path, "a", encoding="utf-8") as f:
            f.write(message + "\n")
    except IOError:
        pass


def prompt_directory() -> str:
    """Prompt user for directory."""
    return input("Enter directory path: ").strip()


def prompt_keyword() -> str:
    """Prompt user for keyword."""
    return input("Enter keyword to filter (optional): ").strip()


def prompt_sort() -> bool:
    """Prompt sort order."""
    ans = input("Sort by size descending? (y/n): ").strip().lower()
    return ans == "y"


def main() -> None:
    """Main CLI entry."""
    log_file = "file_tool.log"
    directory = prompt_directory()
    records = list_files(directory)

    if not records:
        print("No files found.")
        return

    keyword = prompt_keyword()
    if keyword:
        records = filter_files(records, keyword)

    reverse = prompt_sort()
    records = sort_files(records, reverse)

    display(records)

    stats = summarize(records)
    print(f"Files: {stats['count']}, Total Size: {stats['total_size']} bytes")

    action = input("Read a file? (filename or skip): ").strip()
    if action and action != "skip":
        path = os.path.join(directory, action)
        content = read_file(path)
        if content is not None:
            print("--- Content Preview ---")
            print(content[:200])
            append_log(log_file, f"Read file: {path}")

    save = input("Save summary to file? (y/n): ").strip().lower()
    if save == "y":
        output = f"Files: {stats['count']}\nTotal Size: {stats['total_size']}\n"
        if write_file("summary.txt", output):
            print("Saved to summary.txt")
            append_log(log_file, "Saved summary.")

    print("Done.")


if __name__ == "__main__":
    main()
