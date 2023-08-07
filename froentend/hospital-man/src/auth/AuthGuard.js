import React, { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom';


export default function AuthGuard({ Component,role}) {
    const navigate = useNavigate();
    const [isRenderd,setIsRenderd]=useState(false)

    useEffect(() => {
        console.log("here 1")
       let data3 = localStorage.getItem('data');
        let roleName = localStorage.getItem("role")?.replace(/"/g, "");
        setIsRenderd(true); 
        if (!data3) {
            return navigate("/")
        }

        console.log(role.includes(roleName));
        if(!(role.includes(roleName))){
            return navigate("/")
        }
       

    }, [navigate,role])

    if(isRenderd){
        return <Component />;
    }




    // const navigate = useNavigate();
  

    // let data3 = localStorage.getItem('datakey');
    // let roleName = localStorage.getItem("Role")?.replace(/"/g, "");
    // if (!data3) {
    //     return navigate("/")
    // }

    // console.log(role.includes(roleName));
    // if(!(role.includes(roleName))){
    //     return navigate("/")
    // }

    // const navigate = useNavigate();

    // useEffect(() => {
    //     let data3 = localStorage.getItem('datakey');
    //     if(!data3){
    //         navigate("/")
    //     }
    // },[])


    // return <Component />;

}