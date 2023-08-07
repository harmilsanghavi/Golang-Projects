import React, { useState, useEffect } from 'react'
import { useForm } from "react-hook-form";
import axios from 'axios'
import {useNavigate} from 'react-router-dom';

export const Appointments = () => {
  const [preAppoinmentdata, setPreAppoinmentData] = useState();
  const [selectedDepartmentOption, setSelectedDepartmentOption] = useState();
  const [selectedDoctorOption, setSelectedDoctorOption] = useState();
  const [AvailableTime, setAvailableTimeOption] = useState();
  let data3 = localStorage.getItem("token")?.replace(/"/g, "");
  const navigate = useNavigate();
  useEffect(() => {
    // let local
    if (data3) {
      //local = JSON.parse(data3.substring(0, data3.indexOf('}') + 1));

      console.log("first")
      console.log("token in appoinment :- ", data3);
      axios.get(`http://localhost:8900/patient/doctordetail`, {
        headers: {
          Bearer: data3
        }
      })
        .then(response => {
          console.log("pre appoinment data", response.data);
          setPreAppoinmentData(response.data.data);
        })
        .catch(error => {
          console.error(error);
        });
    }
  }, []);
  const { register, handleSubmit } = useForm();
  const onSubmit = async data => {
    //let local = JSON.parse(data3.substring(0, data3.indexOf('}') + 1));
    console.log("form data :- ",data);
    const age=parseInt(data.age)
    const number=parseInt(data.number)
    const doctor=parseInt(data.doctor)
    const newData = {...data, age, number,doctor};

    let responese = await axios.post(`http://localhost:8900/patient/bookappoinment`, newData, {
      headers: {
        Bearer: data3
      }
    })
    console.log("form subbmission data :- ",data);
    console.log("response :-", responese);
    navigate("/")
  }
  const handletime=async (event)=>{
    let value=event.target.value;
    let docotoId=document.getElementById("doctor").value;
    const today = new Date();
  const year = today.getFullYear();
  const month = (today.getMonth() + 1).toString().padStart(2, '0'); // JavaScript months are 0-based
  const day = today.getDate().toString().padStart(2, '0');

  const currentDate = `${year}-${month}-${day}`;
  console.log("cuurrent date :- ",currentDate)
  if (value >= currentDate) {
    console.log("value of day :-",value,"value od doctor id :- ",docotoId)
    
    let resposne=await axios.get(`http://localhost:8900/patient/checkDate/${value}/${docotoId}`,{
      headers:{
        Bearer:data3
      }
    })
    console.log("response of time",resposne)
    console.log("response of time",resposne.data.data)
    setAvailableTimeOption(resposne.data.data)
  } else {
    alert("Selected date is in the Past.");
  }
  }
  return (
    <div>
      <div className='container todo-column'>
        <h1>Appointments</h1>
        <form onSubmit={handleSubmit(onSubmit)}>
          <div className='form-group mb-3'>
            <label className='form-label'>First Name</label>
            <input className='form-control' {...register("fname")} />
          </div>
          <div className='form-group mb-3'>
            <label className='form-label'>Last Name</label>
            <input className='form-control' {...register("lname")} />
          </div>
          <div className='form-group mb-3'>
            <label className='form-label'>Number</label>
            <input className='form-control' {...register("number")} />
          </div>
          <div className='form-group mb-3'>
            <label className='form-label'>Address</label>
            <textarea className='form-control' {...register("address")} />
          </div>
          <div className="form-group mb-3" >
            <label className='form-label' htmlFor="Gender">Gender:</label>
            <input type="radio" value="male" name="gender" {...register("gender")} /> Male
            &nbsp;&nbsp;
            <input type="radio" value="female" name="gender" {...register("gender")} /> Female
          </div>
          <div className='form-group mb-3'>
            <label className='form-label'>Age</label>
            <input className='form-control' {...register("age")} />
          </div>
          <div className="form-group">
            <label className='form-label' htmlFor="role">Choose Your Department:</label>
            {preAppoinmentdata && (
              <div>
                <select
                  className="form-control"
                  {...register("department")}
                  value={selectedDepartmentOption}
                  onChange={(e) => {
                    setSelectedDepartmentOption(e.target.value);
                    setSelectedDoctorOption("");
                  }}
                >
                  <option>Select Department</option>
                  {[...new Set(preAppoinmentdata.map((item) => item.Department))].map(
                    (departmentName, index) => (
                      <option key={index} value={departmentName}>
                        {departmentName}
                      </option>
                    )
                  )}
                </select>
                {selectedDepartmentOption && (
                  <div>
                    <label className='form-label'>Choose Your Doctor:</label>
                    <select id="doctor"
                      className="form-control"
                      {...register("doctor")}
                      value={selectedDoctorOption}
                      onChange={(e) => setSelectedDoctorOption(e.target.value)}
                    >
                      <option>Select Doctor</option>
                      {preAppoinmentdata
                        .filter((item) => item.Department === selectedDepartmentOption)
                        .map((item, index) => (
                          <option key={index} value={item.Id}>
                            {item.Name}
                          </option>
                        ))}
                    </select>
                  </div>
                )}
                
              </div>
            )}
          </div>
          <div className='form-group mb-3'>
            <label className='form-label'>Select Date</label>
            <input className='form-control' type="date" {...register("date")} onChange={handletime} />
          </div>
          <div className="form-group">
            <label className='form-label' htmlFor="role">Choose Your Department:</label>
            {AvailableTime && (
              <select className="form-control" {...register("time")}>
                <option>Select Time</option>
                {AvailableTime.map((item, index) => (
                  <option key={index} value={item.ID}>
                    {item.TimeStr}
                  </option>
                ))}
              </select>
            )}
          </div>
          <input type="submit" />
        </form>
      </div>
    </div>
  )
}