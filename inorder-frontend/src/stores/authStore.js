import { create } from "zustand";
import { roles } from "../utils/const";
import { persist } from "zustand/middleware";


const useAuthStore = create(persist((set) => ({
  username: "",
  authToken: "",
  refreshToken: "",
  role : roles.USER,
  setRole : (role) => set({ role: role }),
  setUsername: (username) => set({ username }),
  setAuthToken: (authToken) => set({ authToken }),
  setRefreshToken: (refreshToken) => set({ refreshToken }),
})));

export default useAuthStore;