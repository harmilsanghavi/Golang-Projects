import axios from 'axios';

export const Refresh = async () => {
    let token=localStorage.getItem('refreshToken')?.replace(/"/g, "");
    let data=await axios.get(`http://localhost:8900/refreshtoken`,{
      headers: { Bearer: token }
    })

    console.log("data from refresh token :- ", data);
    return data
}
