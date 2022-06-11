import { useNavigate } from "react-router-dom";

export const AppBar = ({subfolder}) => {
  const navigate = useNavigate();
  return ( <header className="header header--inside">
              <div className="wrapper">
                <div className="header-left">
                  <a href="/" className="header-back" onClick={(e)=>{e.preventDefault(); navigate(-1)}}><svg className="svg ic-arr-back"><use xlinkHref={subfolder?"../../img/sprite.svg#arr-back":"img/sprite.svg#arr-back"}></use></svg> вернуться</a>
                </div>
                <div className="header-right">
                  <div className="logo">
                    <a href="#"><img src={subfolder?"../../img/logo.svg":"img/logo.svg"} alt="" /></a>
                  </div>
                  <div className="header-phone">
                    <div className="phone">+7 987 654-32-10</div>
                    <a href="#" className="btn-recall">Заказать звонок</a>
                  </div>
                </div>
              </div>
            </header>
  )
}