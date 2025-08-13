import { API_URL } from "@/config";
import axios from "axios";

export async function getAllTags() {
    const response = await axios.get(`${API_URL}/api/items/tags`, {withCredentials: true});
    return response.data.map(tag => tag.name);
}

