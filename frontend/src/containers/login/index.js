import React, { useState } from 'react'
import { useDispatch } from 'react-redux'
import {
    Box,
    Grid,
    TextField,
    Typography,
    Button,
    Dialog,
    DialogTitle,
    DialogContent,
    DialogContentText,
    DialogActions,
    Link,
    CircularProgress,
} from '@material-ui/core'
import FilterDramaOutlinedIcon from '@material-ui/icons/FilterDramaOutlined'
import MuiAlert from '@material-ui/lab/Alert'
import authAPI from 'api/login'
import { req } from 'api'
import { Actions } from 'redux/auth'

function Alert(props) {
    return <MuiAlert elevation={6} variant="filled" {...props} />
}

const initialData = {
    projectName: 'fyp-demo-306511',
    zoneName: 'us-central1-a',
    clusterName: 'demo',
    credPath: '/home/ubuntu/fyp/fyp-demo-sa.json',
}
const initialErr = {
    projectName: false,
    zoneName: false,
    clusterName: false,
    credPath: false,
    aggr: false,
}

function isEmptyOrSpaces(str) {
    return str === null || str.match(/^ *$/) !== null
}

export default function Init(props) {
    const dispatch = useDispatch()
    const [data, setData] = useState(initialData)
    const [err, setErr] = useState(initialErr)

    const [authUrlDialog, setAuthUrlDialog] = useState(false)
    const [authUrl, setAuthUrl] = useState('')

    const [isLoading, setIsLoading] = useState(false)

    const [code, setCode] = useState('')

    const handleAuthUrlClose = () => setAuthUrlDialog(false)
    const handleCodeChange = (e) => {
        setCode(e.target.value)
    }
    const handleVerifyCode = async () => {
        try {
            const project = await req(authAPI.verifyCode({ code: code }))
            setAuthUrlDialog(false)
            dispatch(Actions.login(project))
            setIsLoading(false)
        } catch (err) {
            console.error(err)
        }
    }

    const handleLogin = async () => {
        // check empty value
        let er = false
        for (let k in data) {
            if (isEmptyOrSpaces(data[k])) {
                setErr((prev) => {
                    return { ...prev, [k]: true, aggr: true }
                })
                er = true
            }
        }
        if (er) return
        // api call
        try {
            setIsLoading(true)
            const url = await req(authAPI.authenticate(data))
            // check if no URL
            if (isEmptyOrSpaces(url)) {
                throw new Error('No URL found')
            }
            setAuthUrl(url)
            setAuthUrlDialog(true)
            setIsLoading(false)
        } catch (err) {
            setErr((prev) => {
                return { ...prev, aggr: true }
            })
            console.error(err)
            setIsLoading(false)
        }
    }

    const handleChange = (e, id) => {
        let v = e.target.value
        setData((prev) => {
            return { ...prev, [id]: v }
        })
        if (!isEmptyOrSpaces(v)) {
            setErr((prev) => {
                return { ...prev, [id]: false }
            })
        } else {
            setErr((prev) => {
                return { ...prev, [id]: true }
            })
        }
    }

    return (
        <Box display="block" margin="auto" marginTop="5rem" height={'100%'} width={'60%'}>
            <Grid container direction="column" spacing={3}>
                <Grid item style={{ alignSelf: 'center' }}>
                    <FilterDramaOutlinedIcon style={{ fontSize: '5rem' }} />
                </Grid>
                <Grid item style={{ alignSelf: 'center' }}>
                    <Typography variant="h5">Authenticate Your GCP Cluster</Typography>
                </Grid>
                <Grid item>
                    <TextField
                        fullWidth
                        variant="outlined"
                        name="projectName"
                        label={'Project Name'}
                        value={data.projectName}
                        error={err.projectName}
                        onChange={(e) => handleChange(e, 'projectName')}
                    />
                </Grid>
                <Grid item>
                    <TextField
                        fullWidth
                        variant="outlined"
                        name="zoneName"
                        label={'Zone'}
                        value={data.zoneName}
                        error={err.zoneName}
                        onChange={(e) => handleChange(e, 'zoneName')}
                    />
                </Grid>
                <Grid item>
                    <TextField
                        fullWidth
                        variant="outlined"
                        name="clusterName"
                        label={'Cluster Name'}
                        value={data.clusterName}
                        error={err.clusterName}
                        onChange={(e) => handleChange(e, 'clusterName')}
                    />
                </Grid>
                <Grid item>
                    <TextField
                        fullWidth
                        variant="outlined"
                        name="credPath"
                        label={'Absolute Path of Credential File'}
                        value={data.credPath}
                        error={err.credPath}
                        onChange={(e) => handleChange(e, 'credPath')}
                    />
                </Grid>
                <Grid item>
                    <Button onClick={handleLogin} disabled={isLoading} fullWidth variant="outlined">
                        Authenticate &nbsp; &nbsp;
                        {isLoading ? <CircularProgress size={30} /> : null}
                    </Button>
                </Grid>
                {err.aggr ? (
                    <Grid item>
                        <Alert severity="error">Please enter the correct information!</Alert>
                    </Grid>
                ) : null}
            </Grid>
            <Dialog
                open={authUrlDialog}
                onClose={handleAuthUrlClose}
                aria-labelledby="auth-dialog-title"
                aria-describedby="auth-dialog-description"
                maxWidth={'md'}
                disableBackdropClick
            >
                <DialogTitle id="auth-dialog-title">
                    {'Authenticate your Google account'}
                </DialogTitle>
                <DialogContent>
                    <DialogContentText id="auth-dialog-url">
                        <Typography>Login with the following URL</Typography>
                        <Link href={authUrl} target="_blank">
                            {authUrl.slice(0, 100) + '...'}
                        </Link>
                    </DialogContentText>
                    <DialogContentText id="auth-dialog-vericode">
                        <Typography>And paste the verification code here:</Typography>
                        <TextField
                            multiline
                            rows={8}
                            variant="outlined"
                            fullWidth
                            value={code}
                            onChange={handleCodeChange}
                        ></TextField>
                    </DialogContentText>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleVerifyCode}>Done</Button>
                    <Button onClick={handleAuthUrlClose} autoFocus>
                        Cancel
                    </Button>
                </DialogActions>
            </Dialog>
        </Box>
    )
}
