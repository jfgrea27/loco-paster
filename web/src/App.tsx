import React from "react";
import "./styles/App.css";
import PasteItemList from "./components/PasteItemList";

function App() {
  return (
    <div id="root">
      <div className="title-area nes-input">
        <p className="nes-text is-primary">
          LocoPaster
          <img
            src={require("./media/1up.gif")}
            alt="Computer man"
            className="top-image"
          />
        </p>
        <p className="nes-text is-warning">Paste anything you would like.</p>
        <p className="nes-text is-error">
          Copy it later by clicking on its button.
        </p>
        <p className="nes-text is-success">Delete it when you're done!</p>
      </div>
      <PasteItemList />
      <p className="warning-message blink_me">
        Note: Don't paste anything Bowser could use against you!!
      </p>
    </div>
  );
}

export default App;
