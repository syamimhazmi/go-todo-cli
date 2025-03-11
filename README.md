# Todo CLI App

## Commands

### Adding new todo

```bash
./todo -add buy milk
```

### Mark as done

```bash
./todo -complete=1
```

### Delete a todo

```bash
./todo -delete=2
```

### Show listing of todos

```bash
./todo -list
```

### Sample output

```bash
╔═══╤════════════════════╤═══════╤═════════════════════╤═════════════════════╗
║ # │        Task        │ Done  │    Completed At     │     Created At      ║
╟━━━┼━━━━━━━━━━━━━━━━━━━━┼━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━╢
║ 1 │ buy groceries      │ false │ 01 Jan 01 00:00 UTC │ 11 Mar 25 22:30 WIB ║
║ 2 │ finish up todo app │ false │ 01 Jan 01 00:00 UTC │ 11 Mar 25 22:30 WIB ║
╟━━━┼━━━━━━━━━━━━━━━━━━━━┼━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━╢
║                          You have 2 pending todos                          ║
╚═══╧════════════════════╧═══════╧═════════════════════╧═════════════════════╝
```
