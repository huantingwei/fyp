import React from 'react'
import * as d3 from 'd3'
import styles from './forceGraph.module.css'

const linkLength = 300
const linkWidth = 1
const nodeRadius = 40
const labelFontsize = '12px'

export function runForceGraph(container, linksData, nodesData) {
    const links = linksData.map((d) => Object.assign({}, d))
    const nodes = nodesData.map((d) => Object.assign({}, d))

    const containerRect = container.getBoundingClientRect()
    const height = containerRect.height
    const width = containerRect.width

    // sort links
    links.sort(function (a, b) {
        if (a.source > b.source) {
            return 1
        } else if (a.source < b.source) {
            return -1
        } else {
            if (a.target > b.target) {
                return 1
            }
            if (a.target < b.target) {
                return -1
            } else {
                return 0
            }
        }
    })
    //any links with duplicate source and target get an incremented 'linknum'
    for (var i = 0; i < links.length; i++) {
        if (
            i !== 0 &&
            links[i].source === links[i - 1].source &&
            links[i].target === links[i - 1].target
        ) {
            links[i].linknum = links[i - 1].linknum + 1
        } else {
            links[i].linknum = 1
        }
    }

    const getClass = (d) => {
        switch (d.kind) {
            case 'Service':
                return styles.service
            case 'Pod':
                return styles.pod
            case 'Node':
                return styles.node
            case 'Client':
                return styles.client
            default:
                return styles.pod
        }
    }

    const simulation = d3.forceSimulation(nodes).force(
        'link',
        d3
            .forceLink(links)
            .id((d) => d.id)
            .distance(() => linkLength)
    )

    const svg = d3
        .select(container)
        .append('svg:svg')
        .attr('width', width)
        .attr('height', height)
        .attr('viewBox', [-width / 2, -height / 2, width, height])

    svg.append('svg:defs')
        .selectAll('marker')
        .data(links.map((l) => l.content))
        .enter()
        .append('svg:marker')
        .attr('id', String)
        //  .attr("viewBox", "0 -5 10 10")
        //     .attr('viewBox', [-width / 2, -height / 2, width, height])
        .attr('refX', 30)
        .attr('refY', -4.5)
        .attr('markerWidth', 6)
        .attr('markerHeight', 6)
        .attr('orient', 'auto')
        .append('svg:path')
        .attr('d', 'M0,-5L10,0L0,5')

    const link = svg
        .append('svg:g')
        .selectAll('path')
        .data(links)
        .enter()
        .append('svg:path')
        .attr('class', function (d) {
            return 'link ' + d.content
        })
        .attr('id', function (d, i) {
            return 'linkId_' + i
        })
        .attr('marker-end', function (d) {
            return 'url(#' + d.content + ')'
        })
        // .join('path')
        .attr('fill', 'transparent')
        .attr('stroke', '#999')
        // .attr('stroke-opacity', 0.6)
        .attr('stroke-width', 2)

    var linktext = svg.append('svg:g').selectAll('g.linklabelholder').data(links)

    linktext
        .enter()
        .append('g')
        .attr('class', 'linklabelholder')
        .append('text')
        .attr('class', 'linklabel')
        .style('font-size', '10px')
        .attr('x', '50')
        .attr('y', '-20')
        .attr('text-anchor', 'start')
        .style('fill', '#000')
        .append('textPath')
        .attr('xlink:href', function (d, i) {
            return '#linkId_' + i
        })
        .text(function (d) {
            return d.content
        })

    function clamp(x, lo, hi) {
        return x < lo ? lo : x > hi ? hi : x
    }

    const drag = d3.drag().on('start', dragstart).on('drag', dragged)

    function click(event, d) {
        delete d.fx
        delete d.fy
        d3.select(this).classed('fixed', false)
        simulation.alpha(1).restart()
    }

    function dragstart() {
        d3.select(this).classed('fixed', true)
    }

    function dragged(event, d) {
        d.fx = clamp(event.x, -width / 2, width / 2)
        d.fy = clamp(event.y, -height / 2, height / 2)
        simulation.alpha(1).restart()
    }

    const node = svg
        .append('svg:g')
        .attr('stroke', '#fff')
        .attr('stroke-width', linkWidth)
        .selectAll('circle')
        // rect
        // .selectAll('rect')
        .data(nodes)
        .enter()
        .append('svg:circle')
        // .join('circle')
        .attr('r', nodeRadius)
        // .join('rect')
        // .attr('width', 100)
        // .attr('height', 75)
        .classed('node', true)
        // .classed('fixed', (d) => d.fx !== undefined)
        .classed('fixed', true)
        .attr('class', (d) => `${getClass(d)}`)
        // .attr('fill', nodeColor)
        .call(drag)
        .on('click', click)

    const label = svg
        .append('svg:g')
        .attr('class', 'labels')
        .selectAll('text')
        .data(nodes)
        .enter()
        .append('text')
        .style('white-space', 'pre-line')
        .style('font-size', labelFontsize)
        .attr('text-anchor', 'middle')
        .attr('dominant-baseline', 'central')
        .text((d) => {
            return d.name + ':' + d.content
        })

    simulation.on('tick', () => {
        //update link positions
        // link.attr('x1', (d) => d.source.x)
        //     .attr('y1', (d) => d.source.y)
        //     .attr('x2', (d) => d.target.x)
        //     .attr('y2', (d) => d.target.y)
        link.attr('d', function (d) {
            // var dx = d.target.x - d.source.x,
            //     dy = d.target.y - d.source.y,
            // var dr = 75 / d.linknum //linknum is defined above
            return (
                'M' +
                d.source.x +
                ',' +
                d.source.y +
                'L' +
                // 'A' +
                // dr +
                // ',' +
                // dr +
                // ' 0 0,1 ' +
                d.target.x +
                ',' +
                d.target.y
            )
        })

        // update node positions
        node.attr('cx', (d) => d.x).attr('cy', (d) => d.y)
        // node.attr('x', (d) => d.x).attr('y', (d) => d.y)
        // update label positions
        label
            .attr('x', (d) => {
                return d.x
            })
            .attr('y', (d) => {
                return d.y
            })
    })

    return {
        destroy: () => {
            simulation.stop()
        },
        nodes: () => {
            return svg.node()
        },
    }
}

export function ForceGraph({ linksData, nodesData, nodeHoverTooltip }) {
    const containerRef = React.useRef(null)

    React.useEffect(() => {
        let destroyFn

        if (containerRef.current) {
            const { destroy } = runForceGraph(
                containerRef.current,
                linksData,
                nodesData,
                nodeHoverTooltip
            )
            destroyFn = destroy
        }

        return destroyFn
    }, [linksData, nodeHoverTooltip, nodesData])

    return <div ref={containerRef} className={styles.container} />
}
