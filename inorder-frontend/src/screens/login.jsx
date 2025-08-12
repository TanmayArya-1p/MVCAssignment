import { useEffect, useState } from 'react';
import * as auth from '../api/auth';
import useAuthStore from '../stores/authStore';
import toast, { Toaster } from 'react-hot-toast';
import { jwtDecode } from "jwt-decode";
import { roles } from '../utils/const';
import VerifySignedIn from '../utils/verify';
import { useNavigate } from 'react-router-dom';

export default function LoginScreen() {

    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const {setUsername: setStoreUsername, setAuthToken, setRefreshToken, setRole} = useAuthStore.getState();
    const navigate = useNavigate();
    useEffect(() => {VerifySignedIn()}, [])


    const handleLogin = async (e) => {
        e.preventDefault();
        try {
            let resp = await auth.LoginUser(username, password);
            setStoreUsername(username);
            setAuthToken(resp.authToken);
            setRefreshToken(resp.refreshToken);
            setRole(jwtDecode(resp.authToken).role);
            toast.success("Successfully logged in");
            setTimeout(() => {
                navigate("/home");
            }, 1000);
        } catch(err) {
            console.error("Login failed:", err);
            toast.error("Invalid username or password");
        }
    }


    return (
        <>
        <Toaster />
        <title>Login - InOrder</title>
        <div className="h-screen w-screen flex justify-center items-center">
            <div className="flex flex-col justify-center py-12 lg:px-8">
                <div className="sm:mx-auto sm:w-full sm:max-w-sm">
                    <a href="/">
                        <div className="ubuntu-bold text-4xl w-full text-center">InOrder</div>
                    </a>

                </div>

                <div className="mt-10 card">
                    <h2 className="text-center text-3xl ubuntu-bold tracking-tight">
                        Sign in to your account
                    </h2>
                    <form onSubmit={handleLogin} className="space-y-6 mt-10">
                        <div>
                            <label htmlFor="username" className="block text-sm ubuntu-bold -100">
                                Username
                            </label>
                            <div className="mt-2">
                                <input
                                    id="username"
                                    name="username"
                                    type="username"
                                    value={username}
                                    onChange={(e) => setUsername(e.target.value)}
                                    required
                                    autoComplete="username"
                                    className="w-full bg-white rounded-sm border-2 p-2 "
                                />
                            </div>
                        </div>

                        <div>
                            <div className="flex items-center justify-between">
                                <label htmlFor="password" 
                                className="block text-sm ubuntu-bold -100">
                                    Password
                                </label>
                            </div>
                            <div className="mt-2">
                                <input
                                    id="password"
                                    name="password"
                                    type="password"
                                    value={password}
                                    onChange={(e) => setPassword(e.target.value)}
                                    required
                                    autoComplete="current-password"
                                    className="w-full bg-white rounded-sm border-2 p-2 "
                                />
                            </div>
                        </div>

                        <div>
                            <button
                                type="submit"
                                className="flex w-full justify-center"
                            >
                                Sign in
                            </button>
                        </div>
                    </form>

                    <p className="mt-10 text-center text-sm/6 -400">
                        Don't have an account?{' '}
                        <a href="/register" className="links font-semibold text-indigo-400 hover:text-indigo-300">
                            Sign Up
                        </a>
                    </p>
                </div>
            </div>
        </div>
        </>
    )
}