import { create, createStore } from "zustand";
import { roles } from "../utils/const";


const useAuthStore = create((set) => ({
  username: "",
  authToken: "",
  refreshToken: "",
  roles : roles.USER,
  setUsername: (username) => set({ username }),
  setAuthToken: (authToken) => set({ authToken }),
  setRefreshToken: (refreshToken) => set({ refreshToken }),
}));

export default useAuthStore;