import axios from 'axios';
import { API_URL } from '../config';



export async function RegisterUser(username, password) {
    let data = JSON.stringify({
        "username": username,
        "password": password
    });

    let config = {
    method: 'post',
    maxBodyLength: Infinity,
    url: `${API_URL}/api/auth/register`,
    headers: { 
        'Content-Type': 'application/json'
    },
    data : data,
    withCredentials: true
    };
    try {
        const response = await axios.request(config);
        console.log(JSON.stringify(response.data));
        return response.data;
    } catch (error) {
        console.error(error);
    }

}


export async function LoginUser(username, password) {
    let data = JSON.stringify({
    "username": username,
    "password": password
    });

    let config = {
    method: 'post',
    maxBodyLength: Infinity,
    url: `${API_URL}/api/auth/login`,
    headers: { 
        'Content-Type': 'application/json', 
        'Accept': 'application/json'
    },
    data : data,
    withCredentials: true
    };

    try {
        const response = await axios.request(config);
        console.log(JSON.stringify(response.data));
        return response.data;
    }
    catch (error) {
        throw(error);
    }
}


export async function RefreshToken() {

    let config = {
    method: 'get',
    maxBodyLength: Infinity,
    url: `${API_URL}/api/auth/refresh`,
    withCredentials: true
    };

    try {
        let response = await axios.request(config);
        console.log(JSON.stringify(response.data));
        return response.data;
    }
    catch (error) {
        throw(error);
    }
}


export async function VerifyToken() {

    let config = {
    method: 'get',
    maxBodyLength: Infinity,
    url: `${API_URL}/api/auth/verify`,
    withCredentials: true
    };

    try {
        let response = await axios.request(config);
        console.log(JSON.stringify(response.data));
        return response.data;
    }
    catch (error) {
        throw(error);
    }
}

export async function LogoutUser() {

    let config = {
    method: 'post',
    maxBodyLength: Infinity,
    url: `${API_URL}/api/auth/logout`,
    withCredentials: true
    };

    try {
        let response = await axios.request(config);
        console.log(JSON.stringify(response.data));
        return response.data;
    }
    catch (error) {
        throw(error);
    }
}