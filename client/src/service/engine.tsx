import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:8080",
});

export const collect = async (value: string) => {
  try {
    const response = await api.post(`/collect?value=${value}`);
    console.log("Collection result:", response.data);
    if (response.data && response.data.redirect) {
      window.location.href = response.data.redirect;
    }
    return response.data;
  } catch (error) {
    console.error("Error during collection:", error);
  }
};
