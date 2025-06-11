# Traffic Lights - Exploring State Machines in Go

This project demonstrates how to model a traffic light system using finite state machines (FSMs) in Go. It's a practical example of applying FSM principles to simulate real-world systems.

---

## Project Structure

```
traffic-lights/
├── cmd/
│   └── traffic-light/    # Entry point for the application
├── pkg/
│   └── fsm/              # Core FSM implementation
├── go.mod                # Go module file
└── README.md             # Project documentation
```

---

## Getting Started

### Prerequisites

- Go 1.20 or later

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/dscholtz/traffic-lights.git
   cd traffic-lights
   ```

2. **Build the application:**

   ```bash
   go build -o traffic-light ./cmd/traffic-light
   ```

3. **Run the application:**

   ```bash
   ./traffic-light
   ```

---

## Usage

Upon running the application, you'll see the traffic light cycling through its states:

```
Current State: Red
Current State: Green
Current State: Yellow
...
```

Each state transition occurs after a predefined duration, simulating a real traffic light's behavior.

---

## How It Works

The FSM is defined with three primary states:

- **Red**
- **Green**
- **Yellow**

Transitions between these states are time-based, emulating the standard operation of traffic lights.

The `pkg/fsm` package contains the core logic for the FSM, including state definitions and transition rules.

---

Adjust these values to simulate different traffic scenarios.

---

## Learn More

This project is a part of a broader exploration into state machines in Go. For more examples and detailed explanations, visit the [State Machines in Go](https://github.com/dscholtz/state-machines-in-go) repository.

---

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to enhance the project.

---
