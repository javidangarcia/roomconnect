import { searchResultSchema } from "../schemas";
import { getToken } from "../token";
import type { SearchBody, UserAttributes } from "../types";

export const searchUsers = async (filter: Partial<SearchBody>): Promise<UserAttributes[]> => {
  const url = new URL("/search", import.meta.env.VITE_BACKEND_URL);

  for (const [key, val] of Object.entries(filter)) {
    url.searchParams.set(key, val);
  }

  const response = await fetch(url, {
    headers: {
      Authorization: `Bearer ${getToken()}`,
    },
  });

  const data = await response.json();
  const responseData = searchResultSchema.parse(data);

  if ("errorMessage" in responseData) {
    throw new Error(responseData.errorMessage);
  }

  return responseData;
};