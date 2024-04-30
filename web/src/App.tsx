import React from "react";
import logo from "./logo.svg";
import "./App.css";
import PasteItemList from "./components/PasteItemList";

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <PasteItemList />
      </header>
    </div>
  );
}

export default App;
