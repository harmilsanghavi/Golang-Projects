import React from 'react'
import Carousel from 'react-bootstrap/Carousel'
import '../asset/css/style.css'

export const Home = () => {
  let ReactImg = require('../asset/image/personalized-banner.jpeg')
  let VueImg = require('../asset/image/pioneer-banner.jpeg')
  let AngularImg = require('../asset/image/prevention-banner.jpeg')
  return (
    <div className='container todo-column'>
      <Carousel controls indicators>
        <Carousel.Item>
          <img className="d-block w-100" src={ReactImg} alt="slide 1" />
        </Carousel.Item>
        <Carousel.Item>
          <img className="d-block w-100" src={VueImg} alt="slide 2" />
        </Carousel.Item>
        <Carousel.Item>
          <img className="d-block w-100" src={AngularImg} alt="slide 3" />
        </Carousel.Item>
      </Carousel>


      <section class="section facility-section">

          <div class="row">
            <div class="col-md-4">
              <h2 class="mb-4">Facilities and<br/>Support Services</h2>
              <p class="mb-5">With multidisciplinary health care, specialists from different disciplines work together to diagnose and treat patients accurately.&ZeroWidthSpace;</p>
              <a href="#" class="btn btn-outline-primary d-none d-sm-inline-flex">View All</a>
            </div>
            <div class="col-md-8">
              <div class="row">
                <div class="col-md-3 card">
                  <a href="#" class="text-dark">
                    <div class="text-center facility-box mb-3">
                      <img src="https://www.hcghospitals.in/wp-content/uploads/2022/10/icu-icon-new2.png" alt="HCG Facility" width="30" class="mb-2" />
                        <p class="mb-0">Intensive Care</p>
                    </div>
                  </a>
                </div>
                <div class="col-md-3 card">
                  <a href="#" class="text-dark">
                    <div class="text-center facility-box mb-3">
                      <img src="https://www.hcghospitals.in/wp-content/uploads/2022/11/ambulance-new-icon.png" alt="HCG Facility" width="30" class="mb-2" />
                        <p class="mb-0">Ambulance Service</p>
                    </div>
                  </a>
                </div>
                <div class="col-md-3 card">
                  <a href="#" class="text-dark">
                    <div class="text-center facility-box mb-3">
                      <img src="https://www.hcghospitals.in/wp-content/uploads/2022/12/pathology-icon-new.png" alt="HCG Facility" width="30" class="mb-2"/>
                        <p class="mb-0">Pathology</p>
                    </div>
                  </a>
                </div>
                <div class="col-md-3 card">
                  <a href="#" class="text-dark">
                    <div class="text-center facility-box mb-3">
                      <img src="https://www.hcghospitals.in/wp-content/uploads/2022/12/day-care-icon-new.png" alt="HCG Facility" width="30" class="mb-2"/>
                        <p class="mb-0">Day Care</p>
                    </div>
                  </a>
                </div>
                <div class="col-md-3 card">
                  <a href="#" class="text-dark">
                    <div class="text-center facility-box mb-3">
                      <img src="https://www.hcghospitals.in/wp-content/uploads/2022/12/pharmacy-new-icon.png" alt="HCG Facility" width="30" class="mb-2"/>
                        <p class="mb-0">Pharmacy</p>
                    </div>
                  </a>
                </div>
                <div class="col-md-3 card">
                  <a href="#" class="text-dark">
                    <div class="text-center facility-box mb-3">
                      <img src="https://www.hcghospitals.in/wp-content/uploads/2022/12/radiology-imaging-icon-new.png" alt="HCG Facility" width="30" class="mb-2"/>
                        <p class="mb-0">Radiology and Imaging</p>
                    </div>
                  </a>
                </div>
                <div class="col-md-3 card">
                  <a href="#" class="text-dark">
                    <div class="text-center facility-box mb-3">
                      <img src="https://www.hcghospitals.in/wp-content/uploads/2022/12/microscope-icon-new.png" alt="HCG Facility" width="30" class="mb-2"/>
                        <p class="mb-0">Clinical Research Department</p>
                    </div>
                  </a>
                </div>
                <div class="col-md-3 card">
                  <a href="#" class="text-dark">
                    <div class="text-center facility-box mb-3">
                      <img src="https://www.hcghospitals.in/wp-content/uploads/2022/12/diet-icon-new.png" alt="HCG Facility" width="30" class="mb-2"/>
                        <p class="mb-0">Diet</p>
                    </div>
                  </a>
                </div>
              </div>
              <div class="row mt-4 card">
                <div class="col-md-12 text-center">
                  <a href="#" class="btn btn-outline-primary d-inline-flex d-sm-none">View All</a>
                </div>
              </div>
            </div>
          </div>

      </section>


    </div>
  )
}