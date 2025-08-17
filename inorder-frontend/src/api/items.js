import { API_URL } from "@/config";
import axios from "axios";

export async function getAllItems() {
    const response = await axios.get(`${API_URL}/api/items`, {
        withCredentials: true,
    });
    return response.data;
}

export async function getItemsOfTags(tags) {
    const response = await axios.get(
        `${API_URL}/api/items/bytags?tags=` + tags.join(","),
        { withCredentials: true },
    );
    return response.data;
}

export async function getAllOrderedItems() {
    const response = await axios.get(`${API_URL}/api/orders/items`, {
        withCredentials: true,
    });
    return response.data;
}

export async function bumpOrderItemStatus(itemId) {
    const response = await axios.post(
        `${API_URL}/api/orders/item/${itemId}/bump`,
        {},
        { withCredentials: true },
    );
    return response.data;
}

export async function deleteItem(itemId) {
    const response = await axios.delete(`${API_URL}/api/items/${itemId}`, {
        withCredentials: true,
    });
    return response.data;
}

export async function uploadImage(image) {
    const formData = new FormData();
    formData.append("image", image);
    const response = await axios.post(`${API_URL}/api/items/upload`, formData, {
        headers: {
            "Content-Type": "multipart/form-data",
        },
        withCredentials: true,
    });
    return response.data;
}

export async function createItem({ name, price, description, tags, image }) {
    const response = await axios.post(
        `${API_URL}/api/items`,
        {
            name,
            price: parseFloat(price),
            description,
            tags: tags,
            image,
        },
        { withCredentials: true },
    );
    return response.data;
}

export async function updateItem({
    name,
    price,
    description,
    tags,
    image,
    itemId,
}) {
    const response = await axios.put(
        `${API_URL}/api/items/${itemId}`,
        {
            name,
            price: parseFloat(price),
            description,
            tags: tags,
            image,
        },
        { withCredentials: true },
    );
    return response.data;
}
