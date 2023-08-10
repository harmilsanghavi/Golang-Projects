import React, { useState, useEffect } from 'react'
import axios from 'axios';
import DataTable from 'react-data-table-component';
import jwt from 'jwt-decode'
import { useNavigate } from 'react-router-dom';

const PAGE_SIZE = 5;

export const Dashboard = () => {
  const [tableData, setTableData] = useState([]);
  const [loading, setLoading] = useState(false);
  const [totalRows, setTotalRows] = useState(0);
  const [perPage, setPerPage] = useState(5);
  const [currentPage, setCurrentPage] = useState(1);
  const [sortColumn, setSortColumn] = useState(localStorage.getItem("sortColumn") || "id");
  const [sortDirection, setSortDirection] = useState(localStorage.getItem("sortDirection") || "desc");
  const [searchQuery, setSearchQuery] = useState("");
  const navigate = useNavigate();

  let token = localStorage.getItem("token")?.replace(/"/g, "");
  console.log('token: after', token);

  let role = localStorage.getItem("role")?.replace(/"/g, "");
  console.log("role :- ", role)
  const fetchData = async (page, size = perPage, sort = sortColumn, dir = sortDirection, search = "") => {
    let response
    try {
      let token = localStorage.getItem("token")?.replace(/"/g, "");
      setLoading(true);
      response = await Promise.all([
        role === "patient"
          ? axios.get(`http://localhost:8900/patient/data`, {
            params: {
              current: page,
              size: size,
              column: sort,
              dir: dir,
              search: search,
            },
            headers: { Bearer: token }
          })
          : axios.get(`http://localhost:8900/doctor/data`, {
            params: {
              current: page,
              size: size,
              column: sort,
              dir: dir,
              search: search,
            },
            headers: { Bearer: token }
          })
      ]);

      console.log("response error :- ", response.status)
      console.log("response from dashboard:- ", response[0].data);
      console.log("response from dashboard 2:- ", response[0].data.totalCount);
      setTotalRows(response[0].data.totalCount);
      setTableData(response[0].data.data);
      setLoading(false);
      console.log('totalRows: ', totalRows);
      console.log('currentPage: ', currentPage);
    } catch (error) {

      console.log("response :-", error.response.status);
      if (error.response.status == 401) {
        const data = await refresh();
        localStorage.removeItem('token');
        // localStorage.removeItem('refreshToken');
        localStorage.setItem('token', JSON.stringify(data.data.token));
        // localStorage.setItem('refreshToken', JSON.stringify(data.data.refreshToken));
        fetchData(page)
      }


    }
  };

  const refresh = async () => {
    token = localStorage.getItem('refreshToken')?.replace(/"/g, "");
    const decodedToken = jwt(token);
    console.log('Decoded Token:', decodedToken);
    const expirationTime = new Date(decodedToken.exp * 1000);
    console.log('expirationTime: ', expirationTime);
    const currentTime = new Date();

    if (expirationTime > currentTime) {
      console.log('Token has expired');
      navigate("/")
    }

    let data = await axios.get(`http://localhost:8900/refreshtoken`, {
      headers: { Bearer: token }
    })

    console.log("data from refresh token :- ", data);
    return data
    // localStorage.setItem('token', JSON.stringify(data.data.token));
    // localStorage.setItem('refreshToken', JSON.stringify(data.data.refreshToken));
  }

  useEffect(() => {
    const storedPage = localStorage.getItem("currentPage");
    const initialPage = storedPage ? parseInt(storedPage) : 1;
    setCurrentPage(initialPage);
    fetchData(initialPage);
  }, []);

  const handlePageChange = page => {
    console.log('page: ', page);
    fetchData(page, perPage, sortColumn, sortDirection, searchQuery);
    setCurrentPage(page);
    localStorage.setItem("currentPage", page);
  };
  const handleSort = (column, sortDirection) => {
    console.log('column: ', column.name);
    console.log('sortDirection: ', sortDirection);
    setSortColumn(column.name)
    setSortDirection(sortDirection)
    localStorage.setItem("sortColumn", column.name);
    localStorage.setItem("sortDirection", sortDirection);
    fetchData(currentPage, perPage, sortColumn, sortDirection, searchQuery);
  };

  const handlePerRowsChange = async (newPerPage, page) => {
    fetchData(page, newPerPage);
    setPerPage(newPerPage);
  };
  const handleSearch = async query => {
    setSearchQuery(query);
    fetchData(1, perPage, sortColumn, sortDirection, query);
    setCurrentPage(1);
    localStorage.setItem("currentPage", 1);
  };

  const columns = [
    {
      name: 'Day',
      selector: (row) => row.Day,
      sortable: true,
    },
    {
      name: 'Apooinment_time',
      selector: (row) => row.Time,
      sortable: true,
    },
    {
      name: 'first_name',
      selector: (row) => row.First_name,
      sortable: true,
    },
    {
      name: 'Department',
      selector: (row) => row.Department,
      sortable: true,
    }, {
      name: 'Status',
      selector: (row) => row.Status,
      sortable: true,
    },
    role === "patient" ? (
      {
        name: 'Cancel Appointment',
        // cell: row => <a href={`/cancel/${row.Id}`}>Cancel</a>,
        cell: row => {
          if (row.Day > '2023-08-10') {
            if (row.Status !== 'complete') {
              if (row.Status !== 'reject') {
                return <a href={`/cancel/${row.Id}`}>Cancel</a>
              } else {
                return <p>Rejected</p>
              }
            } else {
              return <p>Completed</p>
            }
          } else {
            return <p>Time Expire</p>
          }
        },
        ignoreRowClick: true,
      }
    ) : (
      {
        name: 'Request',
        cell: row => {
          if (row.Status == 'accept') {
            return <p className='btn btn-sm btn-success' onClick={() => handleUpdateStatus(row, 'complete')}>Clcik To Complete</p>
          } else if (row.Status == 'complete') {
            return <p>Complete</p>
          } else if (row.Status == 'reject') {
            return <p>Rejected</p>
          } else {

            if (row.Day > '2023-08-10') {

              return (<>
                <p className='btn btn-sm btn-primary' onClick={() => handleUpdateStatus(row, 'accept')}>Accept</p>
                &nbsp;&nbsp;
                <p className='btn btn-sm btn-danger' onClick={() => handleUpdateStatus(row, 'reject')}>Reject</p>
              </>)
            } else {
              return <p>Time Expire</p>
            }
          }
        },
        ignoreRowClick: true,
      }
    )
  ];
  const handleUpdateStatus = (row, status) => {
    axios.get(`http://localhost:8900/doctor/appoinmentrequest`, {
      params: {
        Id: row.Id,
        Status: status
      },
      headers: {
        Bearer: token
      }
    })
      .then(response => {
        console.log("response from request", response.data);
        const updatedAppointment = response.data;
        setTableData(prevTableData => {
          const updatedTableData = prevTableData.map(appointment => {
            if (appointment.Id === updatedAppointment.data.ID) {
              return {
                ...appointment,
                Status: updatedAppointment.data.status
              }
            }
            return appointment;
          });
          return updatedTableData;
        });
      })
      .catch(error => {
        console.error(error);
      });
  };

  let canCancelAppointments = role === "patient" ? true : false;
  const columnsToShow = canCancelAppointments
    ? columns.filter(column => column.name !== 'Patient Name')
    : columns.filter(column => column.name !== 'Doctor Name' && column.name !== 'Department');

  return (
    <div className='container todo-column'>
      <h1>Dashboard {role}</h1>
      <div className="search-container">
        <input type="text" placeholder="Search" value={searchQuery} onChange={(e) => handleSearch(e.target.value)} />
      </div>
      {tableData.length > 0 ? (
        <>
          <DataTable
            title="Users"
            columns={columnsToShow}
            data={tableData}
            progressPending={loading}
            pagination
            paginationServer
            paginationPerPage={perPage}
            paginationTotalRows={totalRows}
            paginationDefaultPage={currentPage}
            onChangeRowsPerPage={handlePerRowsChange}
            onChangePage={handlePageChange}
            onSelectedRowsChange={({ selectedRows }) => console.log(selectedRows)}
            onSort={handleSort}
          />
        </>
      ) : (
        <DataTable
          title="Users"
          columns={columnsToShow}
          progressPending={loading}
        />
      )}
    </div>
  )
}
