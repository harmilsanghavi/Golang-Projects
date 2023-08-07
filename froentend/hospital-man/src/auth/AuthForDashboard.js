import React,{useEffect} from 'react'
import { useNavigate } from 'react-router-dom';


export default function AuthForDashboard({ Component }) {
    const navigate = useNavigate();

    useEffect(() => {
        let data3 = localStorage.getItem('data');
        if (data3) {
            return navigate("/dashboard")
        }
    }, [navigate])

    return <Component/>;


}