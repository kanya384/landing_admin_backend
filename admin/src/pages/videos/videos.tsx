import React from "react";

export const Videos: React.FC = () => {
  return (
    <React.Fragment>
      <div className="card-header p-4 border-300 bg-soft">
        <div className="row g-3 justify-content-between align-items-center">
          <div className="col-12 col-md">
            <h1>Видео</h1>
          </div>
          <div className="col col-md-auto">
            <nav className="nav nav-underline justify-content-end border-0 doc-tab-nav align-items-center" role="tablist">
            <button className="btn btn-primary" onClick={() => {}}>Добавить</button>
            </nav>
          </div>
        </div>
      </div>
    </React.Fragment>)
}