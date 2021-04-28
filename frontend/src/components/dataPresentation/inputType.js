import React from 'react'
import uuid from 'react-uuid'
import { makeStyles } from '@material-ui/core/styles'
import PropTypes from 'prop-types'
import {
    Grid,
    TextField,
    Link,
    Typography,
    Chip,
    Accordion,
    AccordionDetails,
    AccordionSummary,
    List,
    ListItem,
} from '@material-ui/core'
import ExpandMoreIcon from '@material-ui/icons/ExpandMore'
import { notNullOrUndefined } from 'utils/transform'

const useAccordionStyles = makeStyles((theme) => ({
    root: {
        width: '100%',
        backgroundColor: '#fafafa',
    },
}))

export default function InputType(props) {
    const accordionClasses = useAccordionStyles()
    const { type, content, primaryKey, secondaryKey } = props
    switch (type) {
        case 'text':
            return <Typography>{content}</Typography>
        case 'multiline':
            return (
                <TextField
                    multiline
                    fullWidth
                    variant="outlined"
                    rows={8}
                    value={content}
                    readOnly
                />
            )
        case 'chip':
            return Object.keys(content).map((v) => {
                return (
                    <div style={{ marginBottom: '0.5rem' }} key={uuid()}>
                        <Chip label={v + ' : ' + content[v]} />
                        {'  '}
                    </div>
                )
            })
        case 'arrayObj': {
            return content.map((val) => {
                return (
                    <Accordion className={accordionClasses.root} key={uuid()} square={true}>
                        <AccordionSummary expandIcon={<ExpandMoreIcon />}>
                            <Typography>
                                {notNullOrUndefined(val) &&
                                notNullOrUndefined(val[primaryKey]) &&
                                Object.keys(val).length !== 0
                                    ? val[primaryKey] +
                                      (notNullOrUndefined(val[secondaryKey])
                                          ? ' : ' + val[secondaryKey]
                                          : '')
                                    : ''}
                            </Typography>
                        </AccordionSummary>
                        <AccordionDetails>
                            <List style={{ width: '100%' }}>
                                {Object.keys(val).map((v) => {
                                    return (
                                        <ListItem key={uuid()}>
                                            {/** object-like */}
                                            {notNullOrUndefined(val[v]) &&
                                            val[v] instanceof Object ? (
                                                <Grid container key={uuid()}>
                                                    <Grid item xs={6}>
                                                        <Typography>{v}</Typography>
                                                    </Grid>
                                                    <Grid item xs={6}>
                                                        {Object.keys(val[v]).length > 0
                                                            ? Object.keys(val[v]).map((obj) => {
                                                                  return (
                                                                      <Chip
                                                                          key={uuid()}
                                                                          label={
                                                                              obj +
                                                                              ' - ' +
                                                                              val[v][obj]
                                                                          }
                                                                      />
                                                                  )
                                                              })
                                                            : null}
                                                    </Grid>
                                                </Grid>
                                            ) : (
                                                <Grid container key={uuid()}>
                                                    <Grid item xs={6}>
                                                        <Typography>{v}</Typography>
                                                    </Grid>
                                                    <Grid item xs={6}>
                                                        <Typography>{val[v]}</Typography>
                                                    </Grid>
                                                </Grid>
                                            )}
                                        </ListItem>
                                    )
                                })}
                            </List>
                        </AccordionDetails>
                    </Accordion>
                )
            })
        }

        case 'link':
            return (
                <Link href={'#'} name={'linkname'} onClick={(e) => console.log(e)}>
                    {content}
                </Link>
            )

        default:
            return null
    }
}

InputType.propTypes = {
    type: PropTypes.string,
    content: PropTypes.any,
    primaryKey: PropTypes.string,
}

InputType.defaultProps = {
    primaryKey: 'name',
    readOnly: true,
}
