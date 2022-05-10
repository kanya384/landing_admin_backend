import { FC } from "react";
import { Doc } from "../../api";

interface TableItemProps {
  doc: Doc,
  index: number,
  editAction: any,
  deleteAction: any
}

export const TableItem: FC<TableItemProps> = (props) => {
  return(<tr>
          <th scope="row">{props.index}</th>
          <td><a href={process.env.REACT_APP_BACKEND_URL!+"/store/"+props.doc.file} target="_blank">{props.doc.title}</a></td>
          <td className="col-sm" style={{width:"235px"}}>
            <div className="">
              <button className="btn btn-soft-success btn-sm" type="button" onClick={()=>{props.editAction()}}>Редактировать</button>
              <button className="btn btn-soft-danger btn-sm" type="button" onClick={()=>{props.deleteAction()}}>Удалить</button>
            </div>
          </td>
        </tr>
  )
}