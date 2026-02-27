# Trying out Bubble Tea 

> "A Go framework based on The Elm Architecture. Bubble Tea is well-suited for simple and complex terminal applications."

<img width="350" height="266" alt="554156183-ad408275-8799-488f-9303-441e7f869535" src="https://github.com/user-attachments/assets/ab2b02aa-52c9-4edc-9f2f-e75f927a01ee" />

I've been fascinated by terminal interfaces for a while 😅. I had only used the super basic ones before, and whenever I tried building my own they were mostly hardcoded and kinda messy.

While looking for easier ways to build TUIs, I found a few frameworks made specifically for this. Textual (Python) and Bubble Tea (Go) kept popping up everywhere as top recommendations.

Since I’ve been learning Go lately, I decided to try Bubble Team, and the whole [Charm](https://github.com/charmbracelet) ecosystem is pretty cool.

---

## Simple demo

1. **Multi-screen Navigation**: Learn how to route users between different "pages" using keyboard numbers.
2. **Progress Bar**: An animated loading bar to see how background timers (ticks) work.
3. **Search & Filter Menu**: A list you can filter in real-time by typing (using the [TextInput component](https://github.com/charmbracelet/bubbles/tree/master/textinput)).
4. **Loading Spinner**: A great way to show users that a process is running in the background without freezing the app.
5. **Help Screen**: A dedicated view explaining navigation and core concepts.

<img width="343" height="222" alt="image" src="https://github.com/user-attachments/assets/e1f1220b-7e24-4b86-bf9e-9a46840badaf" /> <img width="343" height="222" alt="image" src="https://github.com/user-attachments/assets/26ddeb3b-4cda-473c-8f8b-af82e3458dd6" /> <img width="343" height="222" alt="image" src="https://github.com/user-attachments/assets/fd4ecfe0-6a70-407b-aa25-1e493c5a64b1" /> <img width="343" height="222" alt="image" src="https://github.com/user-attachments/assets/72ac60c8-1714-4350-9bfc-f817553d5fb7" />


---

## The Charm Stack

- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)** : The core framework that runs the application loop.
- **[Bubbles](https://github.com/charmbracelet/bubbles)** : Pre-built, reusable UI components like the *spinner*, *text input*, and *progress bar* used in this project.
- **[Lipgloss](https://github.com/charmbracelet/lipgloss)** : The styling engine. It acts like "CSS for your terminal," allowing us to add colors, bold text, borders, and margins.

---

## 🚀 How to Run It!

Make sure you have [Go installed](https://go.dev/doc/install) on your system. 

1. **Clone this repository** (or navigate to the folder).
2. **Install dependencies:**
    ```bash
    go mod tidy
    ```
3. **Run the app!**
    ```bash
    go run main.go

    ```

