# Atlas

**Interactive CLI for API administration with Firebase authentication**

Atlas is a Go-based interactive CLI application that allows users to:

- Load credentials from YAML or manually
- Authenticate with Firebase (`idToken` + `refreshToken`)
- Cache and refresh tokens automatically
- Define API actions in YAML
- Execute API actions with `Authorization: Bearer <idToken>`
- Display JSON responses in table or pretty-printed formats
- Manage multiple environments (dev, QA, prod, custom)

---

## 🚀 Features

- Interactive menu using [survey](https://github.com/AlecAivazis/survey)
- Automatic token caching and refresh
- YAML-defined API endpoints with method, path, and optional body
- JSON array responses shown in **table format**
- Multi-environment support
- Fully interactive CLI workflow

---

## 💻 Installation

1. **Clone the repository**

```bash
git clone https://github.com/your-username/atlas.git
cd atlas
````

2. **Build the CLI**

```bash
go build -o atlas main.go
```

3. **Run the CLI**

```bash
./atlas
```

> On first run, atlas will guide you to load or create credentials and configure your base environment.

---

## ⚙️ Configuration

### Base Config

`~/.config/atlas/config.yaml` (created automatically on first run)

Example:

```yaml
current_env: dev
identity:
  firebase:
    url: "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword"
environments:
  dev:
    base_url: "https://dev.api.example.com"
```

### Credentials

`~/.config/atlas/credentials.yaml` (created manually or saved interactively)

Example:

```yaml
firebase:
  api_key: "YOUR_FIREBASE_API_KEY"
  email: "user@example.com"
  password: "supersecret"
```

### Token Cache

`~/.config/atlas/token.yaml` is automatically generated for session persistence.
Contains `idToken`, `refreshToken`, and expiry.

---

## 📝 Defining API Actions

Create YAML files describing API actions:

```yaml
actions:
  list_users:
    method: GET
    path: /users
    description: "List all users"

  create_user:
    method: POST
    path: /users
    description: "Create a new user"
    body:
      name: "John Doe"
      email: "john@example.com"
```

> Save these in your project and the CLI will load them automatically.

---

## ⚡ Usage

Run the CLI:

```bash
./atlas
```

Interactive menu:

```
What do you want to do?
  ▸ Login (Firebase)
  ▸ Who am I
  ▸ Select environment
  ▸ Call API Actions
  ▸ Exit
```

* Selecting **Call API Actions** will load your YAML-defined actions.
* Prompts appear for missing body fields.
* Array responses are shown as tables automatically.
* Token refresh is handled behind the scenes.

---

## 🔐 Security

* Credentials are saved with file permissions `0600`
* Tokens are cached securely in `~/.config/atlas/token.yaml`
* Token refresh is automatic, no re-login needed unless expired

---

## 🛠️ Dependencies

* [survey](https://github.com/AlecAivazis/survey) – interactive prompts
* [tablewriter](https://github.com/olekukonko/tablewriter) – table display
* [yaml.v3](https://pkg.go.dev/gopkg.in/yaml.v3) – YAML parsing

---

## 📂 Project Structure

```
atlas/
├── main.go
├── cmd/           ← CLI commands and menu
├── auth/          ← Firebase authentication and token cache
├── api/           ← YAML API actions and runner
├── sample/        ← Example API YAML actions
└── README.md
```

---

## ✅ Next Steps

* Add more API actions in YAML
* Flatten nested JSON for table display
* Add multi-environment switching
* Integrate secret manager or keyring

---

## 💡 License

MIT License

