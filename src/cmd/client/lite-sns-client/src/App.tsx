import { Suspense } from "react";
import "./App.css";
import { useRoutes } from "react-router";
import routes from "~react-pages";
import PageBase from "./components/molecules/page_base";

function App() {
  // const [count, setCount] = useState(0)

  return (
    <>
      <Suspense fallback={<p>Loading...</p>}>
        {/* <div>
          <a href="https://vite.dev" target="_blank">
            <img src={viteLogo} className="logo" alt="Vite logo" />
          </a>
          <a href="https://react.dev" target="_blank">
            <img src={reactLogo} className="logo react" alt="React logo" />
          </a>
        </div>
        <h1>Vite + React</h1>
        <div className="card">
          <button onClick={() => setCount((count) => count + 1)}>
            count is {count}
          </button>
          <p>
            Edit <code>src/App.tsx</code> and save to test HMR
          </p>
        </div>
        <p className="read-the-docs">
          Click on the Vite and React logos to learn more
        </p> */}
        <PageBase>{useRoutes(routes)}</PageBase>
      </Suspense>
    </>
  );
}

export default App;
