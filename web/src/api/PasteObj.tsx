import { PasteItemModel, CreatePasteItemModel } from "../models/models";

const buildBaseEndpoint = (): string => {
  if (process.env.NODE_ENV === "development") {
    const port = process.env.LOCO_PASTER_API_PORT || 8000;
    return `http://0.0.0.0:${port}`;
  }
  return "";
};
const commonHeaders = {
  "Access-Control-Allow-Origin": "*",
  "Access-Control-Allow-Methods": "*",
  "Access-Control-Allow-Headers": "*",
};

const apiGetPastes = async (): Promise<PasteItemModel[]> => {
  const endpoint = `${buildBaseEndpoint()}/api/v1/pastes/`;
  const request = new Request(endpoint, {
    method: "GET",
    headers: commonHeaders,
  });

  const response = await fetch(request);
  return await response.json();
};

const apiCreatePaste = async (
  paste: CreatePasteItemModel,
): Promise<PasteItemModel[]> => {
  const endpoint = `${buildBaseEndpoint()}/api/v1/pastes/`;
  const request = new Request(endpoint, {
    method: "POST",
    headers: commonHeaders,
    body: JSON.stringify(paste),
  });

  await fetch(request);

  return await apiGetPastes();
};

const apiDeletePaste = async (id: number): Promise<PasteItemModel[]> => {
  const endpoint = `${buildBaseEndpoint()}/api/v1/pastes/${id}`;
  const request = new Request(endpoint, {
    method: "DELETE",
    headers: commonHeaders,
  });

  await fetch(request);

  return await apiGetPastes();
};

export { apiCreatePaste, apiGetPastes, apiDeletePaste };
