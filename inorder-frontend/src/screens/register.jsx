import { useState, useEffect } from 'react';
import * as auth from '@/api/auth';
import useAuthStore from '@/stores/authStore';
import toast, { Toaster } from 'react-hot-toast';
import VerifySignedIn from '@/utils/verify';
import { useNavigate } from 'react-router-dom';

export default function RegisterScreen() {
    const navigate = useNavigate();
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [repeatPassword, setRepeatPassword] = useState("");
    const {setUsername: setStoreUsername, setAuthToken, setRefreshToken} = useAuthStore.getState();

    useEffect(() => {VerifySignedIn()}, [])
    
    const handleRegister = async (e) => {
        e.preventDefault();
        try {
            if (password !== repeatPassword) {
                toast.error("Passwords do not match");
                return;
            }

            let resp = await auth.RegisterUser(username, password);
            toast.success("Successfully registered");
            setTimeout(() => {
                navigate("/login");
            }, 1000);
        } catch(err) {
            console.error("Registration failed:", err);
            toast.error("User already exists or an error occurred");        
        }
    }


    return (<>
        <Toaster />
        <title>Register - InOrder</title>
        <div className="h-screen w-screen flex justify-center items-center">
            <div className="flex flex-col justify-center py-12 w-[25%] min-w-2xs">
                <div className="sm:mx-auto sm:w-full sm:max-w-sm">
                    <a href='/'>
                        <div className="ubuntu-bold text-4xl w-full text-center">InOrder</div>
                    </a>

                </div>

                <div className="mt-10 sm:mx-auto px-20 w-full card">
                    <h2 className="text-center text-3xl ubuntu-bold tracking-tight">
                        Sign Up
                    </h2>
                    <form onSubmit={handleRegister} className="space-y-6 mt-10">
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
                                    className="w-full"
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
                                    required
                                    value={password}
                                    onChange={(e) => setPassword(e.target.value)}
                                    autoComplete="current-password"
                                    className="w-full"
                                />
                            </div>
                        </div>
                        <div>
                            <div className="flex items-center justify-between">
                                <label htmlFor="password" 
                                className="block text-sm ubuntu-bold -100">
                                    Repeat Password
                                </label>
                            </div>
                            <div className="mt-2">
                                <input
                                    id="repeat-password"
                                    type="password"
                                    required
                                    value={repeatPassword}
                                    onChange={(e) => setRepeatPassword(e.target.value)}
                                    autoComplete="current-password"
                                    className="w-full"
                                />
                            </div>
                        </div>
                        <div>
                            <button
                                type="submit"
                                className="flex w-full justify-center"
                            >
                                Sign Up
                                
                            </button>
                        </div>
                    </form>

                    <p className="mt-10 text-center text-sm/6 -400">
                        Already have an account?{' '}
                        <a href="/login" 
                            className="links font-semibold text-indigo-400 hover:text-indigo-300"
                        >
                            Login
                        </a>
                    </p>
                </div>
            </div>
        </div>
    </>)
}