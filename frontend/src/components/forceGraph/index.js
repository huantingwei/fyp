import React from 'react'
import * as d3 from 'd3'
import styles from './forceGraph.module.css'

const linkLength = 400
const linkWidth = 1
const nodeRadius = 50
const labelFontsize = '12px'
const linkFontsize = '11px'
const legends = ['Service', 'Pod', 'Node', 'External']

function color(d) {
    switch (d) {
        case 'Pod':
            return '#a8bdd3'
        case 'Service':
            return '#fbd491'
        case 'Node':
            return '#63b2e7'
        case 'External':
            return '#899ea8'
        default:
            return '#a8bdd3'
    }
}

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
    // concatenate the content of links with same source-target
    let curr = 0
    for (var i = 1; i < links.length; i++) {
        if (
            links[i].source === links[curr].source &&
            links[i].target === links[curr].target &&
            links[i].content !== links[curr].content
        ) {
            links[curr].content += ' / ' + links[i].content
            links[i].content = ''
        } else {
            curr = i
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
        // .attr('class', 'linklabel')
        .style('font-size', linkFontsize)
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

    // const tooltip = svg.append('svg:g').attr('class', 'tooltip').style('opacity', 0)

    const node = svg
        .append('svg:g')
        .attr('stroke', '#fff')
        .attr('stroke-width', linkWidth)
        .selectAll('circle')
        .data(nodes)
        .enter()
        .append('svg:circle')
        .attr('r', nodeRadius)
        .classed('node', true)
        .classed('fixed', true)
        .attr('class', (d) => `${getClass(d)}`)
        .call(drag)
        .on('click', click)
    // .on('mouseover', function (event, d) {
    //     tooltip.transition().duration(200).style('opacity', 0.9)
    //     tooltip
    //         .html(d.type + '<br/>' + d.name + '<br/>' + d.content)
    //         .style('left', -width / 2 + event.pageX + 'px')
    //         .style('top', -height / 2 + event.pageY - 28 + 'px')
    // })
    // .on('mouseout', function (d) {
    //     tooltip.transition().duration(500).style('opacity', 0)
    // })

    // function wrap(text, width) {
    //     text.each(function () {
    //         var text = d3.select(this),
    //             words = text.text().split(/:+/).reverse(),
    //             word,
    //             line = [],
    //             lineNumber = 0,
    //             lineHeight = 1.1, // ems
    //             x = text.attr('x'),
    //             y = text.attr('y'),
    //             dy = parseFloat(text.attr('dy')),
    //             tspan = text
    //                 .text(null)
    //                 .append('tspan')
    //                 .attr('x', x)
    //                 .attr('y', y)
    //                 .attr('dy', dy + 'em')
    //         while ((word = words.pop())) {
    //             line.push(word)
    //             tspan.text(line.join(' '))
    //             if (tspan.node().getComputedTextLength() > width) {
    //                 line.pop()
    //                 tspan.text(line.join(' '))
    //                 line = [word]
    //                 tspan = text
    //                     .append('tspan')
    //                     .attr('x', x)
    //                     .attr('y', y)
    //                     .attr('dy', ++lineNumber * lineHeight + dy + 'em')
    //                     .text(word)
    //             }
    //         }
    //     })
    // }

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

    const legendLabel = svg
        .append('svg:g')
        .attr('class', 'legend')
        .selectAll('legendLabels')
        .data(legends)
        .enter()
        .append('text')
        .text(function (d) {
            return d
        })
        .attr('text-anchor', 'left')
        .style('alignment-baseline', 'middle')
        .style('font-size', '12px')

    const legendColor = svg
        .append('svg:g')
        .attr('class', 'legend')
        .selectAll('legendColors')
        .data(legends)
        .enter()
        .append('rect')
        .attr('width', 20)
        .attr('height', 20)
        .style('fill', function (d) {
            return color(d)
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
            // var dr = 60 / d.linknum // linknum is defined above
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
        // .call(wrap, 10)

        legendLabel
            .attr('x', (d) => {
                return -width / 2 + 30
            })
            .attr('y', (d, i) => {
                return -height / 2 + 20 + 20 * i
            })
        legendColor
            .attr('x', (d) => {
                return -width / 2 + 5
            })
            .attr('y', (d, i) => {
                return -height / 2 + 10 + 20 * i
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
