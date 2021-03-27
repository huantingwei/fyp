import React from 'react'
import * as d3 from 'd3'
import styles from './forceGraph.module.css'

const linkLength = 400
const linkWidth = 2
const nodeRadius = 40
const labelFontsize = '12px'

export function runForceGraph(container, linksData, nodesData) {
    const links = linksData.map((d) => Object.assign({}, d))
    const nodes = nodesData.map((d) => Object.assign({}, d))

    const containerRect = container.getBoundingClientRect()
    const height = containerRect.height
    const width = containerRect.width

    const getClass = (d) => {
        switch (d.type) {
            case 'Service':
                return styles.service
            case 'Pod':
                return styles.pod
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
        .append('svg')
        .attr('viewBox', [-width / 2, -height / 2, width, height])

    const link = svg
        .append('g')
        .selectAll('line')
        .data(links)
        .join('line')
        .attr('stroke', '#999')
        .attr('stroke-opacity', 0.6)
        .attr('stroke-width', 3)

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
        .append('g')
        .attr('stroke', '#fff')
        .attr('stroke-width', linkWidth)
        .selectAll('circle')
        // rect
        // .selectAll('rect')
        .data(nodes)
        .join('circle')
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
        .append('g')
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
            return d.type + ':' + d.name
        })

    simulation.on('tick', () => {
        //update link positions
        link.attr('x1', (d) => d.source.x)
            .attr('y1', (d) => d.source.y)
            .attr('x2', (d) => d.target.x)
            .attr('y2', (d) => d.target.y)

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
