import axios from 'axios'
import { API_URL } from '../config'

export async function fetchAllUsers() {
    const response = await axios.get(`${API_URL}/api/users` , {withCredentials : true})
    console.log(response.data);
    return response.data;
}

export async function createUser({username,password,role}) {
    const response = await axios.post(`${API_URL}/api/users`, {username,password,role} , {withCredentials: true})
    console.log(response.data);
    return response.data;
}

export async function deleteUser(userID) {
    const response = await axios.delete(`${API_URL}/api/users/${userID}`, {withCredentials: true})
    console.log(response.data);
    return response.data;
}