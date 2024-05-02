import React, { useState, useEffect, ClipboardEvent } from "react";
import PasteItem from "./PasteItem";
import "../styles/PasteItemList.css";

import { PasteItemModel } from "../models/models";
import { apiCreatePaste, apiGetPastes, apiDeletePaste } from "../api/PasteObj";

function PasteItemList() {
  const [text, setText] = useState("");
  const [pasteItems, setPasteItems] = useState<Array<PasteItemModel>>([]);

  useEffect(() => {
    apiFetchPasteItems();
  }, []);

  const apiFetchPasteItems = () => {
    apiGetPastes().then((res) => setPasteItems(res));
  };

  const deletePasteItem = async (id: number) => {
    const deletedItem = await apiDeletePaste(id);
    console.log(deletedItem);
    const remainingItems = pasteItems.filter((e) => e.id != deletedItem.id);
    setPasteItems(remainingItems);
  };

  const addPasteItem = async (event: ClipboardEvent<HTMLInputElement>) => {
    const { clipboardData } = event;
    const text = clipboardData.getData("text");

    const pasteItemModel = {
      blob: text,
    };
    const createdPasteItem = await apiCreatePaste(pasteItemModel);
    const newPasteItems = pasteItems.concat(createdPasteItem);
    setPasteItems(newPasteItems);
  };

  if (pasteItems.length === 0) {
    return (
      <div className="content-area">
        <div className="paste-list-box nes-textarea">
          <div className="grid-container">
            Your pastes will be here..
            <img
              src={require("../media/mario.gif")}
              alt="Computer man"
              className="waiting-image"
            />
          </div>
        </div>

        <div className="paste-area">
          <input
            className="paste-box blink_me nes-textarea"
            type="text"
            value={text}
            onChange={(e) => setText("")}
            placeholder="Paste here"
            onPaste={(e) => addPasteItem(e)}
          />
        </div>
      </div>
    );
  } else {
    return (
      <div className="content-area">
        <div className="paste-list-box nes-textarea">
          <div className="grid-container">
            {pasteItems.map((pasteItem, idx) => (
              <PasteItem
                key={idx}
                id={pasteItem.id}
                blob={pasteItem.blob}
                deletePaste={deletePasteItem}
              />
            ))}
          </div>
        </div>
        <div className="paste-area">
          <input
            className="paste-box blink_me nes-textarea"
            type="text"
            value={text}
            onChange={(e) => setText("")}
            placeholder="Paste here"
            onPaste={(e) => addPasteItem(e)}
          />
        </div>
      </div>
    );
  }
}

export default PasteItemList;
