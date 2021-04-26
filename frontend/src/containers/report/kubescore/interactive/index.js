import React, { useState } from 'react'
import { Button, Grid, Typography } from '@material-ui/core'
import { req } from 'api'
import kubescoreAPI from 'api/kubescore'
import KubeScoreObjList from '../objList'
import StatusHandler from 'components/statusHandler'

const FileUpload = () => {
    const [file, setFile] = useState()
    const [result, setResult] = useState([])
    const [apiStatus, setApiStatus] = useState('initial')
    const [apiMessage, setApiMessage] = useState('')

    const handleChange = (e) => {
        setFile(e.target.files[0])
    }

    const handleSubmit = async () => {
        setApiMessage('')
        setApiStatus('loading')
        let formData = new FormData()
        let uploadedFile = document.querySelector('#file')
        formData.append('file', uploadedFile.files[0])
        try {
            await req(kubescoreAPI._upload(formData))
            setApiStatus('success')
            setResult([])
        } catch (err) {
            console.error(err)
            setApiStatus('fail')
            setApiMessage('Failed to upload file.')
        }
    }

    const handleGet = async () => {
        setApiMessage('')
        setApiStatus('loading')
        try {
            const res = await req(kubescoreAPI._getInteractive())
            if (res === null || res.length === 0) {
                throw new Error('Error')
            }
            // always displays the newest one
            setResult(res['kubescore'])
            setApiStatus('success')
        } catch (err) {
            console.error(err)
            setApiStatus('fail')
            setApiMessage('Failed to perform check.')
        }
    }

    return (
        <Grid container direction="column" spacing={2}>
            <Grid item>
                <Button variant="outlined" component="label">
                    Upload your YAML file
                    <input type="file" id="file" onChange={handleChange} hidden />
                </Button>
                <Typography component="span">
                    &nbsp; &nbsp;{file !== undefined ? file.name : ''}
                </Typography>
            </Grid>
            <Grid item>
                <Button onClick={handleSubmit} variant="outlined" fullWidth>
                    Submit
                </Button>
                <br /> <br />
                <Button onClick={handleGet} variant="outlined" fullWidth>
                    Get Result
                </Button>
            </Grid>
            <Grid item>
                <StatusHandler status={apiStatus} message={apiMessage} minHeight={'3rem'}>
                    {result.length > 0 && result !== undefined ? (
                        <KubeScoreObjList title={'Result'} data={result} />
                    ) : null}
                </StatusHandler>
            </Grid>
        </Grid>
    )
}

export default FileUpload
