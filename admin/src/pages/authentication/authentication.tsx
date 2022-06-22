import React from 'react';
import "./user.min.css"
import { User, Lock } from 'react-feather';
import { useActions } from '../../hooks/use-actions';
import { useState } from 'react';
import { useTypedSelector } from '../../hooks/use-typed-selector';

export const Authentication: React.FC = () => {
    const [form, setForm] = useState({
      login: "",
      pass: ""
    })
    const auths = useTypedSelector(({ auths }) => {
      return auths
    })
    const { sendAuth } = useActions();
    const authClick = (e: any) =>{
      e.preventDefault()
      sendAuth(form.login, form.pass)
    }

    const inputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
      setForm({
        ...form,
        [e.target.name]:e.target.value
      })
    }
    return (
        <main className="main" id="top">
          <div className="container-fluid px-0">
            <div className="container">
              <div className="row flex-center min-vh-100 py-5">
                <div className="col-sm-10 col-md-8 col-lg-5 col-xl-5 col-xxl-3">
                  <a className="d-flex flex-center text-decoration-none mb-4" href="/" onClick={(e)=>{e.preventDefault()}}>
                    <div className="d-flex align-items-center">
                      <img src="assets/img/icons/logo_dark.png" alt="phoenix" width="105" />
                    </div>
                  </a>
                  <div className="text-center mb-7">
                    <h3>Аутентификация</h3>
                    <p className="text-700">Панель управления сайтом</p>
                  </div>
                  {auths?.error!==""?<div className="alert alert-soft-danger">{auths?.error}</div>:<div></div>}
                  <div className="mb-3 text-start">
                      <label className="form-label" htmlFor="email">Логин</label>
                      <div className="form-icon-container">
                          <input className="form-control form-icon-input" onChange={inputChange} id="email" type="text" placeholder="Введите логин"  name="login" value={form.login} />
                          <User size={14} className='form-icon' />
                      </div>
                  </div>
                  <div className="mb-3 text-start">
                      <label className="form-label" htmlFor="password">Пароль</label>
                      <div className="form-icon-container">
                          <input className="form-control form-icon-input" onChange={inputChange} type="password" placeholder="Введите пароль" name="pass" value={form.pass} />
                          <Lock size={14} className='form-icon' />
                      </div>
                  </div>
                  <div className="row flex-between-center mb-7">
                      <div className="col-auto"></div>
                      <div className="col-auto">
                        {/*<div className="form-check mb-0">
                          <input className="form-check-input" id="basic-checkbox" type="checkbox" checked />
                          <label className="form-check-label mb-0" htmlFor="basic-checkbox">Remember me</label>
                        </div>*/}
                      </div>
                  </div>
                  <button className="btn btn-primary w-100 mb-3" onClick={authClick}>Войти</button>
                </div>
              </div>
            </div>
          </div>
        </main>
    );
}