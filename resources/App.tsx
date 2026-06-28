export default function App() {
    return (
        <main className="container">
            <h1>Digo Framework</h1>

            <p>
                Selamat datang di Digo 🚀
            </p>

            <button
                onClick={() => {
                    alert("Hello from Digo!");
                }}
            >
                Click Me
            </button>
        </main>
    );
}