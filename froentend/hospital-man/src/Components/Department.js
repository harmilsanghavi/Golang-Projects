import React, { useEffect, useState } from 'react'
import axios from 'axios'
import "../asset/css/style.css"

export const Department = () => {
  const [data, setData] = useState([]);
  const [isShown,setIsShow]= useState(null)
  console.log('isShow', data)
  useEffect(() => {
    axios.get(`http://localhost:8900/getdepartment`)
      .then(response => {
        setData(response.data);
      })
      .catch(error => {
        console.error(error);
      });
  }, [])

  const handleClick = (id) =>{
    console.log('id', id)
    setIsShow(isShown ===  id ? "null" : id)
  }
  
  

  return (
    <div>
      <div className='container todo-column'>
        <h1>Department</h1>
        <div className='dept'>
        {Array.from(new Set(data.map(dept => dept.Department_name))).map(deptName => {
      const deptDetails = data.filter(dept => dept.Department_name === deptName);
      return (
        <div class="card" onClick={() => handleClick(deptName)}>
          <div class="card-body">
            <h5 class="card-title">{deptName}</h5>
            {isShown === deptName && (
              <div>
                <div>
                  {deptDetails.map(dept => (
                    <div>
                      <p>Dr. {dept.Fname} {dept.Lname}</p>
                    </div>
                  ))}
                </div>
              </div>
            )}
          </div>
        </div>
      );
    })}
        </div>
      </div>
    </div>
  )
}
