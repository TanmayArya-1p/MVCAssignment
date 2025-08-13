import axios from 'axios';
import { API_URL } from '@/config';



export async function RegisterUser(username, password) {
    const response = await axios.post(`${API_URL}/api/auth/register`, {username, password}, {
        withCredentials: true
    });
    return response.data;
}


export async function LoginUser(username, password) {
    const response = await axios.post(`${API_URL}/api/auth/login`, {username, password}, {
        withCredentials: true
    });
    return response.data;
}


export async function RefreshToken() {
    const response = await axios.get(`${API_URL}/api/auth/refresh`, {
        withCredentials: true
    });
    return response.data;
}   

export async function VerifyToken() {
    const response = await axios.get(`${API_URL}/api/auth/verify`, {
        withCredentials: true
    });
    return response.data;
}

export async function LogoutUser() {
    const response = await axios.post(`${API_URL}/api/auth/logout`, {}, {
        withCredentials: true
    });
    return response.data;
}