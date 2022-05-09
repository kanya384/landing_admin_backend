import elements from "./elements"
import { useNavigate, useLocation } from "react-router-dom";
const MenuComponent = () => {
  const location = useLocation().pathname;
  let navigate = useNavigate();
  const pushHistory = (url: string) => {
    navigate(url, { replace: true })
  }
  return(
        <nav className="navbar navbar-light navbar-vertical navbar-vibrant navbar-expand-lg">
          <div className="navbarVerticalCollapse">
          <div className="navbar-logo justify-content-center">
            <button className="btn navbar-toggler navbar-toggler-humburger-icon" type="button" data-bs-toggle="collapse" data-bs-target="#navbarVerticalCollapse" aria-controls="navbarVerticalCollapse" aria-expanded="false" aria-label="Toggle Navigation">
              <span className="navbar-toggle-icon">
                <span className="toggle-line"></span>
              </span>
            </button>
            <a className="navbar-brand me-1 me-sm-3" href="../../index.html">
              <div className="d-flex justify-content-center">
                  <img src="../../assets/img/icons/logo.png" alt="phoenix" width="105" />
              </div>
            </a>
          </div>
          <div className="navbar-vertical-content scrollbar">
            <ul className="navbar-nav flex-column">
              {elements.map((element)=>{
                if (element.divider) {
                  return <p key={element.name} className="navbar-vertical-label">{element.name}</p>
                }
                return  <li key={element.url} className="nav-item">
                          <a className={location === element.url?"nav-link active":"nav-link"} onClick={(e)=>{e.preventDefault(); pushHistory(element.url!)}} href={element.url!}>
                            <div className="d-flex align-items-center">
                              <span className="nav-link-icon">{element.icon}</span>
                              <span className="nav-link-text">{element.name}</span>
                            </div>
                          </a>
                        </li>
              })}
            </ul>
          </div>
        </div>
      </nav>
  )
}

export default MenuComponent