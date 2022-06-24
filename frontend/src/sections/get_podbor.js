import { EditableText } from "../components/editable-text"
import { RangeSlider } from "../components/range-slider";
import React, {useContext, useEffect, useState} from "react";
import {Link, useLocation} from "react-router-dom";
import {ContentContext} from "../context/contentContext";
import { Form } from "../components/form";
import { Modal } from "../components/modals";

export const GetPodbor = () => {
  const location = useLocation();
  const content = useContext(ContentContext)
  const [isOpen, setIsOpen] = useState(false)

  const [filter, setFilter] = useState({
    areaMin: "",
    areaMax: "",
    floors: [1,9],
    rooms: [1,2,3],
    liters: [1,2,3],
    sort: "area",
  })
  
  const setFloors = (value) => {
    setFilter({
      ...filter,
      floors: value
    })
  }
  
  const setRooms = (event) => {
    let rooms = parseInt(event.currentTarget.getAttribute("rooms"))
    if (event.target.checked && !filter.rooms.includes(rooms)) {
      setFilter({
        ...filter,
        rooms: [
          ...filter.rooms,
          rooms
        ]
      })
    } else {
      setFilter({
        ...filter,
        rooms: filter.rooms.filter(room => room!=rooms)
      })
    }
  }
  
  const setLiters = (event) => {
    let liter = parseInt(event.currentTarget.getAttribute("liter"))
    if (event.target.checked && !filter.liters.includes(liter)) {
      setFilter({
        ...filter,
        liters: [
          ...filter.liters,
          liter
        ]
      })
    } else {
      setFilter({
        ...filter,
        liters: filter.liters.filter(liters => liter!=liters)
      })
    }
  }
  
  const resetFilter = (e) => {
    setFilter({
      areaMin: "",
      areaMax: "",
      floors: [1,9],
      rooms: [1,2,3],
      liters: [1,2,3],
    })
  }
  
  const sortFlats = (filedName, flats) => {
    
    return flats.sort((a,b)=>{
      if (a[filedName] > b[filedName]) {
        return 1;
      }
      if (a[filedName] < b[filedName]) {
        return -1;
      }
      return 0;
    })
  }
  
  useEffect(() => {
    window.scrollTo(0, 0);
  }, [location]);
  
  
  return(
    <div className="lvl12" style={{zIndex: 900000}}>
      <div className="wrapper">
        <div className="podborka-left">
          <div className="title"><EditableText id={"62aef61ba26e626025a8d8da"} defaultText={"Получите подборку свободных квартир с планировками и ценами"}/></div>
          <div className="bottom-text">
            <EditableText id={"62aef61ba26e626025a8d8db"} defaultText={"Перезвоним и предоставим самые выгодные предложения на сегодня"}/>
          </div>
        </div>
        <div className="podborka-right">
          <div className="calculator-section">
            <div className="calculator-filter">
              <div className="filter-item filter-item__floor">
                <div className="filter-item__title">Этаж</div>
                <div className="filter-item__body">
                  <RangeSlider value={filter.floors} setValue={(floor) => setFloors(floor)} min={1} max={9} />
                </div>
              </div>
              <div className="filter-item filter-item__rooms">
                <div className="filter-item__body">
                  <ul className="checked-list checked-list--column-3">
                    <li>
                      <label className="inp-checkbox"><input id="ck1" type="checkbox" rooms={1} onChange={setRooms} checked={filter.rooms.includes(1)} />
                        <svg className="svg checked"><use xlinkHref="img/sprite.svg#checked"></use></svg>
                        <span className="label">1-комн.</span></label>
                    </li>
                    <li>
                      <label className="inp-checkbox"><input id="ck2" type="checkbox" liter={1} onChange={setLiters} checked={filter.liters.includes(1)} />
                        <svg className="svg checked"><use xlinkHref="img/sprite.svg#checked"></use></svg>
                        <span className="label">Корпус 1</span></label>
                    </li>
                    <li>
                      <label className="inp-checkbox"><input type="checkbox" rooms={2} onChange={setRooms} checked={filter.rooms.includes(2)} />
                        <svg className="svg checked"><use xlinkHref="img/sprite.svg#checked"></use></svg>
                        <span className="label">2-комн.</span></label>
                    </li>
                    <li>
                      <label className="inp-checkbox"><input type="checkbox" liter={2} onChange={setLiters} checked={filter.liters.includes(2)}/>
                        <svg className="svg checked"><use xlinkHref="img/sprite.svg#checked"></use></svg>
                        <span className="label">Корпус 2</span></label>
                    </li>
                    <li>
                      <label className="inp-checkbox"><input type="checkbox" rooms={3} onChange={setRooms} checked={filter.rooms.includes(3)} />
                        <svg className="svg checked"><use xlinkHref="img/sprite.svg#checked"></use></svg>
                        <span className="label">3-комн.</span></label>
                    </li>
                    <li>
                      <label className="inp-checkbox"><input type="checkbox" liter={3} onChange={setLiters} checked={filter.liters.includes(3)} />
                        <svg className="svg checked"><use xlinkHref="img/sprite.svg#checked"></use></svg>
                        <span className="label">Корпус 3</span></label>
                    </li>
                  </ul>
                </div>
              </div>
              <div className="filter-item">
                <div className="form-ec form-ec--calculator">
                  <div className="form-ec__content">
                    <div className="form-ec__input-row">
                      <Form fields={[
                            {
                              type:"text",
                              name: "name",
                              placeholder: "Имя",
                              required: false,
                            },
                            {
                                type:"text",
                                name: "phone",
                                placeholder: "Телефон",
                                required: true,
                            },
                            {
                              type:"text",
                              name: "email",
                              placeholder: "E-mail",
                              required: true,
                            }, 
                          ]} 
                          btnTitle={"Получить подборку"} 
                          description={`Получить подборку квартир. Этаж:${filter.floors.join("-")}; Кол-во комнат:${filter.rooms.join(",")}; Корпус:${filter.liters.join(",")}`}
                          celtype={"getPodbor"}
                          close={()=>{}} 
                          callback={()=>{setIsOpen(true)}}
                      />
                    </div>
                    
                  </div>
                </div>
              </div>
            </div>
          </div>
          
        </div>
      </div>
      <div className="wrapper">
        <div className="b-link-row">
            <Link to="/promo" className="lnk-params js-open-modal-flat">
              <span className="lnk-params__ico">
              <svg className="ico" width="66" height="30" viewBox="0 0 66 30"
                    xmlns="http://www.w3.org/2000/svg">
                <path d="M64.5 14.9983L65.2033 15.7092L65.9219 14.9982L65.2032 14.2874L64.5 14.9983ZM51.0509 29.7109L65.2033 15.7092L63.7967 14.2875L49.6443 28.2891L51.0509 29.7109ZM65.2032 14.2874L51.0508 0.289034L49.6444 1.71097L63.7968 15.7093L65.2032 14.2874ZM64.5 13.9983H0.5V15.9983H64.5V13.9983Z" />
              </svg>
              </span>
              <EditableText id={"62af012c2f3ab6f9b4a854dc"} defaultText={"И еще множество причин жить на юге"}/>
            </Link>
        </div>
      </div>
      <Modal 
          success={true}
          position={window.pageYOffset}
          opened={isOpen}
          close = {()=>{setIsOpen(null)}}
        />
    </div>
  )
}

export default GetPodbor