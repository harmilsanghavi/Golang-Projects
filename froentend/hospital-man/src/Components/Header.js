import React from 'react'
import { Outlet, Link, useNavigate } from "react-router-dom";
import { Modal } from "react-bootstrap";
import { useState } from "react";
import { useForm } from "react-hook-form";
import axios from 'axios'
// import Spinner from './Spinner'


export const Header = ({ handleLogin, handleLogout }) => {
    // const [isLoading, setIsLoading] = useState(false);
    const [modalState, setModalState] = useState("close");
    const [selectedOption, setSelectedOption] = useState("");
    const navigate = useNavigate();


    const isAuthenticated = localStorage.getItem('isAuthenticated') === 'true';

    let role = localStorage.getItem('role')
    let roleboolen = false
    if (role) {
        let str = role;
        str = str.replace(/"/g, "");
        console.log(str);
        console.log("role in froentend", role)
        if (str == "patient") {
            roleboolen = true
        } else {
            roleboolen = false
        }
    }

    const handleShowModalOne = () => {
        setModalState("modal-one")
    }

    const handleShowModalTwo = () => {
        setModalState("modal-two")
    }

    const handleClose = () => {
        setModalState("close")
    }
    const { register, handleSubmit, formState: { errors } } = useForm();

    const onSubmit = (data) => {
        const { role_name, age, number, gender, first_name, qualification, department, last_name, email, password } = data;
        const userData = {
            email,
            password,
            role_name,
        };
        console.log("userData :- ", userData)
        let doctorData = {};

        if (role_name === "doctor") {
            doctorData = {
                first_name,
                last_name,
                age: parseInt(age),
                number: parseInt(number),
                gender,
                qualification,
                department,
            };
            console.log("data in formate of json:::", data)
        } else {
            doctorData = {}
        }
        data = JSON.stringify({ user: userData, doctor: doctorData });


        axios.post('http://localhost:8900/signup', data,
            { headers: { 'Content-Type': 'application/json' } }
        )
            .then(response => { console.log("sing up response :-", response.data) })
            .catch(error => { console.log(error.data) });
        console.log(data)
        handleClose()
        navigate("/")

    };

    const onFormSubmit = (e) => {
        e.preventDefault(); // prevent the default form submission behavior
        handleSubmit(onSubmit)(e); // call the custom submit function instead
    }
    const [formdata, setFormData] = useState({ email: "", password: "" });

    const handlechange = (e) => {
        const { name, value } = e.target;
        setFormData((prevFormdata) => ({ ...prevFormdata, [name]: value }));

    };
    const handleSignIn = async (e) => {
        e.preventDefault();
        // localStorage.setItem("data",JSON.stringify(formdata))
        var data = JSON.stringify(formdata)
        // console.log("data to be subimmted :- ", data)
        handleClose()
        // setIsLoading(true);
        let responese = await axios.post(`http://localhost:8900/signin`, data)
        console.log("sign in data:- ", responese);
        console.log("sign of data to be:- ", responese.data.data);
        console.log("Token:- ", responese.data.token);
        console.log("role:- ", responese.data.role);
        console.log("ID:- ", responese.data.id);
        // setIsLoading(false);
        if (responese.data.status === false) {
            console.log("false")
            handleShowModalTwo()
            navigate("/")
        } else {
            console.log("true")
            localStorage.setItem('token', JSON.stringify(responese.data.token));
            localStorage.setItem('refreshToken', JSON.stringify(responese.data.refreshToken));
            localStorage.setItem('id', JSON.stringify(responese.data.id));
            localStorage.setItem('role', JSON.stringify(responese.data.role));
            localStorage.setItem('data', JSON.stringify(responese.data.data));
            handleLogin();
            // setIsLoggedIn(true)
            navigate("/dashboard")
        }
    }

    const ClearStorage = () => {
        const logoutConfirmed = window.confirm('Are you sure you want to log out?');

        if (logoutConfirmed) {
            localStorage.removeItem('token');
            localStorage.removeItem('id');
            localStorage.removeItem('role');
            localStorage.removeItem('data');
            localStorage.removeItem('currentPage');
            localStorage.removeItem('sortColumn');
            localStorage.removeItem('sortDirection');
            localStorage.removeItem('refreshToken');
            roleboolen = false
            // setIsLoggedIn(false)
            handleLogout();
            navigate("/")
        }
    }


    return (
        <div>

            <nav className="navbar navbar-expand-lg bg-body-tertiary ">
                <div className="container-fluid">
                    <a className="navbar-brand" href="#"><p><i className="fa-solid fa-circle-h"></i></p></a>
                    <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
                        <span className="navbar-toggler-icon"></span>
                    </button>
                    <div className="collapse navbar-collapse" id="navbarSupportedContent">
                        <ul className="navbar-nav me-auto mb-2 mb-lg-0">
                            <li className="nav-item">
                                <Link className="nav-link active" aria-current="page" to="/">Home</Link>
                            </li>
                            <li className="nav-item">
                                <Link className="nav-link" to="/departments">Departments</Link>
                            </li>
                            <li className="nav-item">
                                {roleboolen &&
                                    <Link className="nav-link" to="/appointments">Appointments</Link>
                                }
                            </li>
                            <li className="nav-item">
                                <Link className="nav-link" to="/about">About us</Link>
                            </li>
                        </ul>
                        <Outlet />
                        <div className="d-flex " role="search">
                            {isAuthenticated ?
                                <div>
                                    <p className='btn btn-outline-danger' onClick={ClearStorage}>Logout</p>
                                </div>
                                :
                                <div> <p className='btn btn-primary' variant="primary" onClick={handleShowModalOne}>Sign Up</p>
                                    &nbsp; &nbsp;
                                    <p className='btn btn-primary' variant="primary" onClick={handleShowModalTwo}>Sign In</p>
                                </div>
                            }
                        </div>
                    </div>
                </div>
            </nav>
            {/* model sign up bootstrap */}
            <Modal show={modalState === "modal-one"}>
                <Modal.Header>
                    <Modal.Title>Sign Up </Modal.Title>
                    <p onClick={handleClose}><i className="fa-solid fa-circle-xmark"></i></p>
                </Modal.Header>
                <Modal.Body>
                    <form method='post' onSubmit={onFormSubmit} className='sign-up-form'>
                        <div className="form-group">
                            <label className='form-label' htmlFor="email">Email:</label>
                            <input className='form-control'
                                {...register("email", { required: true })
                                } />
                            {errors.email && <p>Please check the Email</p>}
                        </div>
                        <div className="form-group">
                            <label className='form-label' htmlFor="role">Choose Your Role:</label>
                            <select className='form-control' {...register("role_name")} value={selectedOption} onChange={(e) => setSelectedOption(e.target.value)} >
                                <option>Select Role</option>
                                <option value="doctor">Doctor</option>
                                <option value="patient">Patient</option>
                                <option value="staff">Staff</option>
                            </select>
                        </div>
                        {selectedOption === "doctor" &&
                            <div className="mt-2">
                                <div className="form-group mb-3">
                                    <label className='form-label' htmlFor="FirstName">FirstName:</label>
                                    <input className='form-control' {...register("first_name")} />
                                </div>
                                <div className="form-group mb-3">
                                    <label className='form-label' htmlFor="LastName">LastName:</label>
                                    <input className='form-control' {...register("last_name")} />
                                </div>
                                <div className="form-group mb-3" >
                                    <label className='form-label' htmlFor="Gender">Gender:</label>
                                    <input type="radio" value="male" name="gender" {...register("gender")} /> Male
                                    &nbsp;&nbsp;
                                    <input type="radio" value="female" name="gender" {...register("gender")} /> Female
                                </div>
                                <div className='form-group mb-3'>
                                    <label className='form-label' htmlFor="Age">Age:</label>
                                    <input className='form-control' {...register("age")} />
                                </div>
                                <div className='form-group mb-3'>
                                    <label className='form-label' htmlFor="Address">Qualification:</label>
                                    <input className='form-control' {...register("qualification")} />
                                </div>
                                <div className='form-group mb-3'>
                                    <label className='form-label' htmlFor="department">Department:</label>
                                    <input className='form-control' {...register("department")} />
                                </div>
                                <div className='form-group mb-3'>
                                    <label className='form-label' htmlFor="mobile">Mobile No:</label>
                                    <input className='form-control' {...register("number")} />
                                </div>
                            </div>}
                        <div className="form-group mb-3">
                            <label className='form-label' htmlFor="password">Password:</label>
                            <input className='form-control' type="password" {...register("password")} />
                        </div>
                        <div className="form-group mb-3">
                            <input className='form-control btn btn-info ' type='submit' />
                        </div>
                    </form>
                </Modal.Body>
            </Modal>


            {/* sign in model */}

            <Modal show={modalState === "modal-two"}>
                <Modal.Header>
                    <Modal.Title>Sign In</Modal.Title>
                    <p onClick={handleClose}><i className="fa-solid fa-circle-xmark"></i></p>
                </Modal.Header>
                <Modal.Body>
                    <form onSubmit={handleSignIn}>
                        <div className="form-group">
                            <label className='form-label' htmlFor="email">Email:</label>
                            <input className='form-control' value={formdata.email} onChange={handlechange} name="email" />
                        </div>
                        <div className="form-group mb-3">
                            <label className='form-label' htmlFor="password">Password:</label>
                            <input className='form-control' type="password" value={formdata.password} onChange={handlechange} name="password" />
                        </div>
                        <div className="form-group mb-3">
                            <input className='form-control btn btn-info ' type='submit' />
                        </div>
                    </form>
                </Modal.Body>
            </Modal>
            {/* <Spinner loading={isLoading} /> */}
        </div>
    )
}



