import Navbar from "../components/navbar"
import VerifySignedIn from "../utils/verify";
import { useEffect } from "react";
export default function LandingScreen() {

    useEffect(() => {VerifySignedIn()}, [])
    //TODO: ADD LOADING TILL VERIFIED

    return <>
        <div className="h-screen w-screen flex flex-col justify-center items-center">
            <h1 className="text-3xl ubuntu-bold text-center mt-10">
                InOrder
            </h1>
            <div className="mt-5 flex flex-row p-2 gap-2">
                <button>
                    <a href="/login" className="ubuntu-regular text-black">
                        Login
                    </a>
                </button>
                <button>
                    <a href="/register" className="ubuntu-regular text-black">
                        Register
                    </a>
                </button>
            </div>



        </div>
    
    
    
    </>
}
