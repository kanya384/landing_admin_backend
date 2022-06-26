import { title } from "process";
import { FC } from "react";
import { Title } from "../../api";

interface TableItemProps {
  title: Title,
  index: number,
  editAction: any,
  deleteAction: any
}

export const TableItem: FC<TableItemProps> = (props) => {
  return(<tr>
          <th scope="row">{props.index}</th>
          <td>{props.title.desktopTitle}</td>
          <td>{props.title.mobileTitle}</td>
          <td>{props.title.tagName}</td>
          <td>{props.title.tagValue}</td>
          <td className="col-sm" style={{width:"235px"}}>
            <div className="">
              <button className="btn btn-soft-success btn-sm" type="button" onClick={()=>{props.editAction()}}>Редактировать</button>
              <button className="btn btn-soft-danger btn-sm" type="button" onClick={()=>{props.deleteAction()}}>Удалить</button>
            </div>
          </td>
        </tr>
  )
}