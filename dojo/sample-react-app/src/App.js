import "./App.css";
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";

function App() {
  function Home() {
    return <h1>Home</h1>;
  }

  function About() {
    return <h1>About</h1>;
  }

  function Blog() {
    return <h1>Blog</h1>;
  }

  function Layout() {
    return (
      <div>
        <nav>
          <ul>
            <li>
              <Link to="/">Home</Link>
            </li>
            <li>
              <Link to="/about">About</Link>
            </li>
            <li>
              <Link to="/blog">Dashboard</Link>
            </li>
          </ul>
        </nav>
      </div>
    );
  }

  return (
    <div>
      <Router>
        <Routes>
          <Route index element={<Home />}></Route>
          <Route path="/blog" element={<Blog />}></Route>
          <Route path="/about" element={<About />}></Route>
        </Routes>
        <Layout />
      </Router>
    </div>
  );
}

export default App;
