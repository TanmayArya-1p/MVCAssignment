import { useNavigate } from "react-router-dom";
import Spinner from "../components/spinner";
import VerifySignedIn from "../utils/verify";
import { useEffect, useState } from "react";



export default function LandingScreen() {
    const navigate = useNavigate();
    const[loading,setLoading] = useState(true);


    useEffect(() => {    
        setLoading(true);
        VerifySignedIn().then(() => setLoading(true)).catch(() => setLoading(false));

    }, [])
    return <>
        <title>InOrder</title>
        <div className="h-screen w-screen flex flex-col justify-center items-center">
            {loading ? <Spinner /> : <>
                <h1 className="text-3xl ubuntu-bold text-center mt-10">
                    InOrder
                </h1>
                <div className="mt-5 flex flex-row p-2 gap-2">
                    <button className="cursor-pointer" onClick={() => navigate("/login")}>
                        Login
                    </button>
                    <button className="cursor-pointer" onClick={() => navigate("/register")}>
                        Register
                    </button>
                </div>
                </>
            }


        </div>
    
    
    
    </>
}
