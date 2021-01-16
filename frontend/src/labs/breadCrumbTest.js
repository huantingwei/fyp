// import React, { Fragment, useEffect } from 'react'
// import { useSelector, useDispatch } from 'react-redux'
// import { Breadcrumbs, Button } from '@material-ui/core'
// import KubeBenchReportList from '../containers/report/kubebench/list'
// import KubeBenchSection from '../containers/report/kubebench/section'
// import KubeBenchDetail from '../containers/report/kubebench/detail'
// import { Actions } from 'redux/kubebench'

// export default function KubeBench() {
//     const state = useSelector((state) => state.kubebench)
//     console.log('state:', state)

//     const items = [
//         {
//             title: 'CIS Report List',
//             content: <KubeBenchReportList />,
//         },
//         {
//             content: <KubeBenchSection />,
//         },
//         {
//             content: <KubeBenchDetail />,
//         },
//     ]

//     const handleClick = (index) => {
//         console.log(index)
//         Actions.changeLevel({ level: index + 1 })
//     }

//     return (
//         <Fragment>
//             <Breadcrumbs aria-label="breadcrumb">
//                 {/* initiate array of length = state.level */}
//                 {Array(state.level)
//                     .fill(0)
//                     .map((_, i) => {
//                         return (
//                             <Button onClick={() => handleClick(i)} key={i}>
//                                 {items[i]['title']}
//                             </Button>
//                         )
//                     })}
//             </Breadcrumbs>
//             {items[state.level - 1].content}
//         </Fragment>
//     )
// }
