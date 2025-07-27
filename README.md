# qMuxr

A blazing-fast, local search multiplexer for your browser's default search. Navigate the web with custom bangs, shortcuts, and session launchers, all powered by a tiny, self-hosted Go service.

## Features

*   **Bang Search (`!keyword`):** Directly jump to your favorite search engines or websites with custom prefixes (e.g., `!g go modules`, `!yt cat videos`).
*   **Web Shortcuts (`#shortcut_name`):** Create personalized, quick links to any URL (e.g., `#portal` to open your school portal).
*   **Session Launcher (`@session_name`):** Open multiple predefined tabs with a single command (e.g., `@dev` to open GitHub, Stack Overflow, and your dev server).
*   **Local & Private:** Your configurations are stored locally in a SQLite database, never leaving your machine.
*   **Blazing Fast:** Runs as a tiny background service, providing instant redirects without any network latency or browser pop-up blockers.
*   **Cross-Platform:** Built with Go, available for Linux, macOS, and Windows.

## Get Started

### 1. Installation

**Recommended (Coming Soon: Package Managers):**
*   **Linux (APT/Debian):** `sudo apt install qmuxr`
*   **Linux (Pacman/Arch):** `sudo pacman -S qmuxr`
*   **macOS (Homebrew):** `brew install qmuxr`
*   **Windows (Scoop/Chocolatey):** `choco install qmuxr` (or `scoop install qmuxr`)

**Manual Installation (for now):**
1.  **Download:** Grab the latest executable from the [Releases page](https://github.com/your-org/qmuxr-cli/releases) for your operating system.
2.  **Place it:** Move the `qmuxr` (or `qmuxr.exe` on Windows) executable to a directory in your system's `PATH` (e.g., `/usr/local/bin` on Linux/macOS, or `C:\Program Files\qMuxr` then add to PATH on Windows).
3.  **Permissions (Linux/macOS):** Make it executable: `chmod +x /path/to/qmuxr`.

### 2. Start the Local Server

Run the `qmuxr` HTTP server in the background:

```bash
qmuxr serve
```

**For persistent background service (Recommended):**
*   **Linux (systemd):** Instructions for setting up a systemd service unit will go here (e.g., `sudo systemctl enable --now qmuxr`).
*   **macOS (launchd):** Instructions for setting up a launchd agent.
*   **Windows (Task Scheduler / NSSM):** Instructions for running as a background task.

The server will listen on `http://localhost:5980` by default.

### 3. Configure Your Browser

Set `http://localhost:5980/?q=%s` as your browser's default search engine.

**Instructions for common browsers:**
*   **Zen:** `Settings > Search > Search Shortcuts`
    *  Don't forget to set qMuxr as your default search engine at the top of the page. 
*   **Chromium:** `Settings > Search engine > Manage search engines and site search > Add`
    *   **Search engine:** `qMuxr`
    *   **Shortcut:** `qmx` (or anything you prefer)
    *   **URL with %s:** `http://localhost:5980/?q=%s`
*   **Firefox:** `Settings > Search > Add new search engine`
    *   Instructions for adding a custom engine might be slightly more involved, or you might need a dedicated extension.

---

## Usage

Once configured, simply type into your browser's address bar:

*   **Bang Search:**
    *   `!g my Go query`
    *   `!yt epic jdm cars drifting`
    *   `!gh qmuxr`
*   **Web Shortcut:**
    *   `#portal`
    *   `#book`
*   **Session Launcher:**
    *   `@dev`
    *   `@learning`

---

## CLI Commands

`qmuxr` offers commands to manage your bangs, shortcuts, and sessions:

*   **`qmuxr serve`**: Starts the local HTTP server.
*   **`qmuxr add-bang <prefix> <url_template>`**: Add or update a bang.
    *   Example: `qmuxr add-bang so https://stackoverflow.com/search?q=%s`
*   **`qmuxr add-shortcut <name> <url>`**: Add or update a web shortcut.
    *   Example: `qmuxr add-shortcut portal https://myschool-portal.com/student-dashboard`
*   **`qmuxr add-session <name> <url1> [url2...]`**: Add or update a session.
    *   Example: `qmuxr add-session dev https://github.com/user https://stackoverflow.com https://your-local-dev-server:3000`
*   **`qmuxr list`**: (Coming Soon) List all configured bangs, shortcuts, and sessions.
*   **`qmuxr delete <type> <name/prefix>`**: (Coming Soon) Delete a bang, shortcut, or session.
    *   Example: `qmuxr delete bang old`
    *   Example: `qmuxr delete shortcut old-link`
    *   Example: `qmuxr delete session old-session`

---

## Ideas
- [ ] "Shortcuts" e.g. repos => github.com/johndoe?tab=repositories
- [ ] Analytics
- [ ] Settings dashboard or cli
- [x] History

## Features 

## Shortcuts
- [ ] Cache
- [x] Handler function
- [x] Service
- [x] Repository
- [x] Parse them seperately from bangs
- [x] Handler struct
- [x] Use interfaces for dependency injection
- [x] Determine precedence order

### History
- [x] Repository
- [x] Service
- [x] Go routine for db transactions
- [ ] Middleware

## Todo
- [x] Setup database connection
- [x] Setup database migrations
- [x] scrape & dump duckduckgo's bang index into the relational db
- [x] Implement query & bang parsing 
- [x] Check bang matches against db and keep highest ranking one 
- [x] Implement service layer 
- [x] Implement url resolving
- [x] fix cors
- [x] Implement provider fallback
- [x] Speed it up
- [x] clean up api & router
- [x] implement caching using sync.Map
- [ ] Implement features: shortcuts #, sessions @ and history $
- [ ] build cli
- [ ] Embed snapshot of the database & hydrate a local version on the user's machine
