import React, { useState, useEffect, ClipboardEvent } from "react";
import PasteItem from "./PasteItem";

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
    const fetchedPasteItems = await apiDeletePaste(id);
    setPasteItems(fetchedPasteItems);
  };

  const addPasteItem = async (event: ClipboardEvent<HTMLInputElement>) => {
    const { clipboardData } = event;
    const text = clipboardData.getData("text");

    const pasteItemModel = {
      blob: text,
    };
    const fetchedPasteItems = await apiCreatePaste(pasteItemModel);
    setPasteItems(fetchedPasteItems);
  };

  return (
    <div className="paste-item-list">
      {pasteItems.map((pasteItem, idx) => (
        <PasteItem
          key={idx}
          id={pasteItem.id}
          blob={pasteItem.blob}
          deletePaste={deletePasteItem}
        />
      ))}
      <input
        type="text"
        value={text}
        onChange={(e) => setText("")}
        placeholder="Paste something here"
        onPaste={(e) => addPasteItem(e)}
      />
    </div>
  );
}

export default PasteItemList;
