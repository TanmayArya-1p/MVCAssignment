import Spinner from "@/components/spinner";
import VerifySignedIn from "@/utils/verify";
import { useEffect, useState } from "react";

export default function SignedIn({children}) {
    const [loading, setLoading] = useState(true)
    
    useEffect(() => {
        setLoading(true);
        VerifySignedIn().then(() => setLoading(false)).catch(() => {
            setLoading(true);
        });
    },[])


    if(loading) {
        return <div className="h-screen w-screen flex items-center justify-center">
            <Spinner/>
        </div>
    }

    return <>
        {children}
    </>
}