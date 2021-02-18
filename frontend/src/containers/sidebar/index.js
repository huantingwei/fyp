import React, { Fragment, useState } from 'react'
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
    Collapse,
    MenuItem,
    Menu,
} from '@material-ui/core'
import IconButton from '@material-ui/core/IconButton'
import MenuIcon from '@material-ui/icons/Menu'
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft'
import ChevronRightIcon from '@material-ui/icons/ChevronRight'
import CloudQueueOutlinedIcon from '@material-ui/icons/CloudQueueOutlined'
import AccountBoxOutlinedIcon from '@material-ui/icons/AccountBoxOutlined'
import { ExpandLess, ExpandMore } from '@material-ui/icons'
import { Link as RouterLink, useLocation } from 'react-router-dom'

import { Actions } from 'redux/auth'
import { useDispatch } from 'react-redux'

const drawerWidth = 260

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
    profile: {
        position: 'absolute',
        right: '3rem',
    },
}))

export default function LeftDrawer(props) {
    const { listItems, botItems, children } = props
    const dispatch = useDispatch()
    const classes = useStyles()
    const theme = useTheme()
    const [pathName, setPathName] = useState(useLocation().pathname)
    const [open, setOpen] = useState(true)
    const [moduleOpen, setModuleOpen] = useState(
        listItems.reduce((acc, curr) => {
            if (curr.nested) {
                return { ...acc, [curr.id]: false }
            } else {
                return acc
            }
        }, {})
    )
    const [profileAnchor, setProfileAnchor] = React.useState(null)

    const handleProfileClick = (event) => {
        setProfileAnchor(event.currentTarget)
    }

    const handleProfileClose = () => {
        setProfileAnchor(null)
    }
    const handleDrawerOpen = () => {
        setOpen(true)
    }

    const handleDrawerClose = () => {
        setOpen(false)
    }

    const handleItemClick = (path) => {
        setPathName(path)
    }

    const handleExpandClick = (index) => {
        setModuleOpen((prev) => {
            return { ...prev, [index]: !moduleOpen[index] }
        })
    }

    return (
        <div className={classes.root}>
            <CssBaseline />
            {/*----------------top bar -------------------- */}
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
                    <div className={classes.profile}>
                        <IconButton
                            color="inherit"
                            aria-label="profile"
                            onClick={handleProfileClick}
                        >
                            <AccountBoxOutlinedIcon />
                            <Typography component="span">&nbsp;&nbsp;Profile</Typography>
                        </IconButton>

                        <Menu
                            id="profile-menu"
                            anchorEl={profileAnchor}
                            keepMounted
                            open={Boolean(profileAnchor)}
                            onClose={handleProfileClose}
                        >
                            <MenuItem onClick={() => dispatch(Actions.logout())}>Logout</MenuItem>
                        </Menu>
                    </div>
                </Toolbar>
            </AppBar>
            {/*----------------top bar -------------------- */}
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
                        {theme.direction === 'rtl' ? <ChevronRightIcon /> : <ChevronLeftIcon />}
                    </IconButton>
                </div>
                <Divider />

                {/** ----------------------drawer---------------------- **/}
                <List>
                    {listItems.map((item, index) => {
                        const { id, icon, text, nested, path } = item
                        return nested ? (
                            /** ----------------------nested---------------------- **/
                            <Fragment key={id}>
                                <ListItem
                                    button
                                    onClick={() => handleExpandClick(id)}
                                    selected={false}
                                >
                                    <ListItemIcon>{icon}</ListItemIcon>
                                    <ListItemText primary={text} />
                                    {moduleOpen[id] ? <ExpandLess /> : <ExpandMore />}
                                </ListItem>
                                <Collapse in={moduleOpen[id]} timeout="auto" unmountOnExit>
                                    <List component="div" disablePadding>
                                        {nested.map((nestedItem) => {
                                            let { id, icon, text, path } = nestedItem
                                            return (
                                                <ListItem
                                                    key={id}
                                                    button
                                                    component={RouterLink}
                                                    to={path}
                                                    onClick={() => handleItemClick(path)}
                                                    selected={path === pathName}
                                                >
                                                    <ListItemIcon>{icon}</ListItemIcon>
                                                    <ListItemText primary={text} />
                                                </ListItem>
                                            )
                                        })}
                                    </List>
                                </Collapse>
                            </Fragment>
                        ) : (
                            /** ----------------------not nested---------------------- **/
                            <ListItem
                                key={id}
                                button
                                component={RouterLink}
                                to={path}
                                selected={path === pathName}
                                onClick={() => handleItemClick(path)}
                            >
                                <ListItemIcon>{icon}</ListItemIcon>
                                <ListItemText primary={text} />
                            </ListItem>
                        )
                    })}
                    <Divider />
                    {/** -------------------bottom items------------------- **/}
                    {botItems.map((item, index) => {
                        const { id, icon, text, path } = item
                        return (
                            <ListItem
                                key={id}
                                button
                                component={RouterLink}
                                to={path}
                                selected={path === pathName}
                                onClick={() => handleItemClick(path)}
                            >
                                <ListItemIcon>{icon}</ListItemIcon>
                                <ListItemText primary={text} />
                            </ListItem>
                        )
                    })}
                </List>
                {/** -------------------bottom items------------------- **/}
                {/** ----------------------drawer---------------------- **/}
            </Drawer>
            {/** ----------------------content---------------------- **/}
            <main className={classes.content}>
                <div className={classes.toolbar} />
                {/* <Typography variant="h4" className={classes.contentTitle}>
                    {pathName}
                </Typography> */}
                {children}
            </main>
            {/** ----------------------content---------------------- **/}
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
            id: 'workload',
            text: 'Workload',
            icon: <CloudQueueOutlinedIcon />,
            nested: [
                {
                    id: 'pod',
                    text: 'Pod',
                    icon: <CloudQueueOutlinedIcon />,
                    content: <Typography>Pod</Typography>,
                },
                {
                    id: 'deployment',
                    text: 'Deployment',
                    icon: <CloudQueueOutlinedIcon />,
                    content: <Typography>Deployment</Typography>,
                },
            ],
        },
    ],
}
