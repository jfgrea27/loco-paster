import React from "react";

type Props = {
  id: number;
  blob: string;
  deletePaste: (id: number) => void;
};

function PasteItem({ id, blob, deletePaste }: Props) {
  return (
    <div className="paste-item">
      <button onClick={() => navigator.clipboard.writeText(blob)}>
        {blob}
      </button>
      <button onClick={() => deletePaste(id)}>X</button>
    </div>
  );
}
export default PasteItem;
