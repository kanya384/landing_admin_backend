import { EditableText } from "../components/editable-text"
import { RangeSlider } from "../components/range-slider";
import React, {useContext, useEffect, useState} from "react";
import {useLocation} from "react-router-dom";
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
                          btnTitle={"Отправить"} 
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
      <Modal 
          success={true}
          position={window.pageYOffset}
          opened={isOpen}
          close = {()=>{setIsOpen(null)}}
        />
    </div>
  )
}
