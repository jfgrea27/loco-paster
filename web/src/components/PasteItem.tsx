import React from "react";
import "../styles/PasteItem.css";

type Props = {
  id: number;
  blob: string;
  deletePaste: (id: number) => void;
};

function PasteItem({ id, blob, deletePaste }: Props) {
  return (
    <div className="grid-item nes-container is-rounded">
      <button
        className="copy-item nes-btn"
        onClick={() => navigator.clipboard.writeText(blob)}
      >
        {blob}
      </button>

      <button
        className="cancel-item nes-btn is-warning"
        onClick={() => deletePaste(id)}
      >
        X
      </button>
    </div>
  );
}

export default PasteItem;
