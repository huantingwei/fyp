import React, { useState } from 'react'
import PropTypes from 'prop-types'
import clsx from 'clsx'
import { makeStyles, useTheme } from '@material-ui/core/styles'
import {
    Drawer,
    AppBar,
    Toolbar,
    List,
    CssBaseline,
    Divider,
    ListItem,
    ListItemIcon,
    ListItemText,
    Typography,
} from '@material-ui/core'
import IconButton from '@material-ui/core/IconButton'
import MenuIcon from '@material-ui/icons/Menu'
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft'
import ChevronRightIcon from '@material-ui/icons/ChevronRight'
import CloudQueueOutlinedIcon from '@material-ui/icons/CloudQueueOutlined'

const drawerWidth = 240

const useStyles = makeStyles((theme) => ({
    root: {
        display: 'flex',
    },
    appBar: {
        zIndex: theme.zIndex.drawer + 1,
        transition: theme.transitions.create(['width', 'margin'], {
            easing: theme.transitions.easing.sharp,
            duration: theme.transitions.duration.leavingScreen,
        }),
        backgroundColor: '#3c414a',
        height: '3rem',
    },
    appBarShift: {
        marginLeft: drawerWidth,
        width: `calc(100% - ${drawerWidth}px)`,
        transition: theme.transitions.create(['width', 'margin'], {
            easing: theme.transitions.easing.sharp,
            duration: theme.transitions.duration.enteringScreen,
        }),
    },
    menuButton: {
        marginRight: 36,
    },
    hide: {
        display: 'none',
    },
    drawer: {
        width: drawerWidth,
        flexShrink: 0,
        whiteSpace: 'nowrap',
    },
    drawerOpen: {
        width: drawerWidth,
        transition: theme.transitions.create('width', {
            easing: theme.transitions.easing.sharp,
            duration: theme.transitions.duration.enteringScreen,
        }),
    },
    drawerClose: {
        transition: theme.transitions.create('width', {
            easing: theme.transitions.easing.sharp,
            duration: theme.transitions.duration.leavingScreen,
        }),
        overflowX: 'hidden',
        width: theme.spacing(7) + 1,
        [theme.breakpoints.up('sm')]: {
            width: theme.spacing(9) + 1,
        },
    },
    toolbar: {
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'flex-end',
        padding: theme.spacing(0, 1),
        height: '3rem',
        maxHeight: '3rem',
        minHeight: '3rem',
        // necessary for content to be below app bar
        // ...theme.mixins.toolbar,
    },
    content: {
        flexGrow: 1,
        padding: theme.spacing(3),
    },
    contentTitle: {
        marginLeft: theme.spacing(2),
        marginBottom: theme.spacing(3),
    },
}))

export default function LeftDrawer(props) {
    const { listItems } = props

    const classes = useStyles()
    const theme = useTheme()
    const [open, setOpen] = useState(false)
    const [selected, setSelected] = useState(listItems[0])

    const handleDrawerOpen = () => {
        setOpen(true)
    }

    const handleDrawerClose = () => {
        setOpen(false)
    }

    const handleItemClick = (index) => {
        setSelected(listItems[index])
    }

    return (
        <div className={classes.root}>
            <CssBaseline />
            <AppBar
                position="fixed"
                className={clsx(classes.appBar, {
                    [classes.appBarShift]: open,
                })}
            >
                <Toolbar style={{ minHeight: '3rem' }}>
                    <IconButton
                        color="inherit"
                        aria-label="open drawer"
                        onClick={handleDrawerOpen}
                        edge="start"
                        size="small"
                        className={clsx(classes.menuButton, {
                            [classes.hide]: open,
                        })}
                    >
                        <MenuIcon />
                    </IconButton>
                    {/* <Typography variant="h6" noWrap>
                        Mini variant drawer
                    </Typography> */}
                </Toolbar>
            </AppBar>
            <Drawer
                variant="permanent"
                className={clsx(classes.drawer, {
                    [classes.drawerOpen]: open,
                    [classes.drawerClose]: !open,
                })}
                classes={{
                    paper: clsx({
                        [classes.drawerOpen]: open,
                        [classes.drawerClose]: !open,
                    }),
                }}
            >
                {/** open/close icon **/}
                <div className={classes.toolbar}>
                    <IconButton onClick={handleDrawerClose}>
                        {theme.direction === 'rtl' ? (
                            <ChevronRightIcon />
                        ) : (
                            <ChevronLeftIcon />
                        )}
                    </IconButton>
                </div>
                <Divider />

                {/** ----------------------drawer content---------------------- **/}
                <List>
                    {listItems.map((item, index) => {
                        const { id, icon, text } = item
                        return (
                            <ListItem
                                button
                                key={id}
                                onClick={() => handleItemClick(index)}
                            >
                                <ListItemIcon>{icon}</ListItemIcon>
                                <ListItemText primary={text} />
                            </ListItem>
                        )
                    })}
                </List>
                <Divider />
                {/** ----------------------drawer content---------------------- **/}
            </Drawer>
            <main className={classes.content}>
                <div className={classes.toolbar} />
                <Typography variant="h4" className={classes.contentTitle}>
                    {selected.text}
                </Typography>
                {selected.content}
            </main>
        </div>
    )
}

LeftDrawer.propTypes = {
    listItems: PropTypes.arrayOf(PropTypes.object),
}

LeftDrawer.defaultProps = {
    listItems: [
        {
            id: 'cluster',
            text: 'Cluster',
            icon: <CloudQueueOutlinedIcon />,
            content: <Typography>Cluster</Typography>,
        },
        {
            id: 'pod',
            text: 'Pod',
            icon: <CloudQueueOutlinedIcon />,
            content: <Typography>Pod</Typography>,
        },
    ],
}
