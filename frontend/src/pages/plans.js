import React, { useEffect } from "react"
import { useLocation, useNavigate } from "react-router-dom";
import { AppBar } from "../components/appbar"
import { EditableText } from "../components/editable-text";
import { Footer } from "../components/footer"

export const Plans = () => {
  const location = useLocation();
  useEffect(() => {
    window.scrollTo(0, 0);
  }, [location]);
  return( <React.Fragment>
      <AppBar />
      <div className="wrapper">
          <div className="page-title"><h1 className="h1"><EditableText id={"62aefdb322de07ce5e31751e"} defaultText={"выбрать квартиру"}/></h1></div>
          <div className="content-wrapper">
            <aside className="left-column">
              <div className="filter-left-btn-toggle">
                <div className="btn-toggle-filter js-btn-toggle-filter">Развернуть параметры</div>
              </div>
              <div className="filter-left js-filter-left-toggle">
                <div className="filter-item">
                  <div className="filter-item__title">Площадь, м²</div>
                  <div className="filter-item__body">
                    <div className="input-group-x2">
                      <div className="input-group-x2__col">
                        <div className="inp-group">
                          <div className="inp-group-label">От</div>
                          <input className="input" placeholder="От" />
                        </div>
                      </div>
                      <div className="input-group-x2__col">
                        <div className="inp-group">
                          <div className="inp-group-label">До</div>
                          <input className="input" placeholder="До" />
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div className="filter-item">
                  <div className="filter-item__title">Этаж</div>
                  <div className="filter-item__body"></div>
                </div>
                <div className="filter-item">
                  <div className="filter-item__body">
                    <ul className="checked-list">
                      <li>
                        <label className="inp-checkbox"><input id="ck1" type="checkbox" />
                          <svg className="svg checked"><use xlinkHref="img/sprite.svg#checked"></use></svg>
                          <span className="label">1-комн.</span></label>
                      </li>
                      <li>
                        <label className="inp-checkbox"><input id="ck2" type="checkbox" />
                          <svg className="svg checked"><use xlinkHref="img/sprite.svg#checked"></use></svg>
                          <span className="label">Корпус 1</span></label>
                      </li>
                      <li>
                        <label className="inp-checkbox"><input type="checkbox" />
                          <svg className="svg checked"><use xlinkHref="img/sprite.svg#checked"></use></svg>
                          <span className="label">2-комн.</span></label>
                      </li>
                      <li>
                        <label className="inp-checkbox"><input type="checkbox" />
                          <svg className="svg checked"><use xlinkHref="img/sprite.svg#checked"></use></svg>
                          <span className="label">Корпус 2</span></label>
                      </li>
                      <li>
                        <label className="inp-checkbox"><input type="checkbox" />
                          <svg className="svg checked"><use xlinkHref="img/sprite.svg#checked"></use></svg>
                          <span className="label">3-комн.</span></label>
                      </li>
                      <li>
                        <label className="inp-checkbox"><input type="checkbox" />
                          <svg className="svg checked"><use xlinkHref="img/sprite.svg#checked"></use></svg>
                          <span className="label">Корпус 3</span></label>
                      </li>
                    </ul>
                  </div>
                </div>
                <div className="filter-btns">
                  <button className="btn-reset">
                    <svg className="svg ico-reset"><use xlinkHref="img/sprite.svg#close"></use></svg> Сбросить
                  </button>
                </div>
              </div>
            </aside>
            <main className="content">
              <div className="listing-filter">
                <div className="listing-sort">
                  <div className="listing-sort__title">Сортировать</div>
                  <div className="listing-sort__select custom-select js-custom-select-wrapper">
                    <div className="custom-select__label js-custom-select-btn-toggle">по площади</div>
                    <div className="custom-select__arr js-custom-select-btn-toggle">
                      <svg className="svg ico-select-arr"><use xlinkHref="img/sprite.svg#select-arr"></use></svg>
                    </div>
                    <div className="custom-select__dropdown js-custom-select-dropdown">
                      <ul className="custom-select__list">
                        <li className="custom-select__item active">по площади</li>
                        <li className="custom-select__item">по цене</li>
                        <li className="custom-select__item">по этажу</li>
                      </ul>
                    </div>
                  </div>
                </div>
              </div>
              <div className="listing-choise">
                <div className="choise-item">
                  <div className="choise-item__img"><img src="img/floor.png" alt="" /></div>
                  <div className="choise-item__right">
                    <div className="choise-item__header">
                      <div className="choise-item__title"><a href="#">1-комн, 43 м²</a></div>
                      <div className="choise-item__floor">1 этаж из 9</div>
                    </div>
                    <div className="choise-item__bottom">
                      <div className="choise-item__link-pdf"><a href="#" className="link-pdf">Скачать в PDF</a></div>
                      <div className="choise-item__more"><a href="#" className="link-bb">узнать стоимость</a></div>
                    </div>
                  </div>
                </div>
                <div className="choise-item">
                  <div className="choise-item__img"><img src="img/floor.png" alt="" /></div>
                  <div className="choise-item__right">
                    <div className="choise-item__header">
                      <div className="choise-item__title"><a href="#">1-комн, 43 м²</a></div>
                      <div className="choise-item__floor">1 этаж из 9</div>
                    </div>
                    <div className="choise-item__bottom">
                      <div className="choise-item__link-pdf"><a href="#" className="link-pdf">Скачать в PDF</a></div>
                      <div className="choise-item__more"><a href="#" className="link-bb">узнать стоимость</a></div>
                    </div>
                  </div>
                </div>
                <div className="choise-item">
                  <div className="choise-item__img"><img src="img/floor.png" alt="" /></div>
                  <div className="choise-item__right">
                    <div className="choise-item__header">
                      <div className="choise-item__title"><a href="#">1-комн, 43 м²</a></div>
                      <div className="choise-item__floor">1 этаж из 9</div>
                    </div>
                    <div className="choise-item__bottom">
                      <div className="choise-item__link-pdf"><a href="#" className="link-pdf">Скачать в PDF</a></div>
                      <div className="choise-item__more"><a href="#" className="link-bb">узнать стоимость</a></div>
                    </div>
                  </div>
                </div>
                <div className="choise-item">
                  <div className="choise-item__img"><img src="img/floor.png" alt="" /></div>
                  <div className="choise-item__right">
                    <div className="choise-item__header">
                      <div className="choise-item__title"><a href="#">1-комн, 43 м²</a></div>
                      <div className="choise-item__floor">1 этаж из 9</div>
                    </div>
                    <div className="choise-item__bottom">
                      <div className="choise-item__link-pdf"><a href="#" className="link-pdf">Скачать в PDF</a></div>
                      <div className="choise-item__more"><a href="#" className="link-bb">узнать стоимость</a></div>
                    </div>
                  </div>
                </div>
                <div className="choise-item">
                  <div className="choise-item__img"><img src="img/floor.png" alt="" /></div>
                  <div className="choise-item__right">
                    <div className="choise-item__header">
                      <div className="choise-item__title"><a href="#">1-комн, 43 м²</a></div>
                      <div className="choise-item__floor">1 этаж из 9</div>
                    </div>
                    <div className="choise-item__bottom">
                      <div className="choise-item__link-pdf"><a href="#" className="link-pdf">Скачать в PDF</a></div>
                      <div className="choise-item__more"><a href="#" className="link-bb">узнать стоимость</a></div>
                    </div>
                  </div>
                </div>
              </div>
            </main>
          </div>
        </div>
        <Footer />
    </React.Fragment>
  )
}