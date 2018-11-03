import React, {
    useState, useEffect, useRef, useContext,
} from 'react';
import './App.scss';
import { Vector as VectorLayer } from 'ol/layer';
import 'ol/ol.css';
import { Map, View, Feature } from 'ol';
import { Point } from 'ol/geom';
import TileLayer from 'ol/layer/Tile';
import { OSM, Vector as VectorSource } from 'ol/source';
import {
    Circle as CircleStyle, Fill, Icon, Stroke, Style,
} from 'ol/style';
import {routeCoords} from './polyLine';

export default () => <App />;

function createMap(elem) {
    return new Map({
        target: elem,
        layers: [
            new TileLayer({
                source: new OSM(),
            }),
        ],
        view: new View({
            center: [-11718716.28195593, 4869217.172379018],
            zoom: 13,
        }),
        maxTilesLoading: 48,
        loadTilesWhileAnimating: true,
        loadTilesWhileInteracting: true,
    });
}

const OLMap = React.createContext(null);
function App() {
    const [state, setState] = useState(null);
    return (
        <OLMap.Provider 
        value={{ 
            map: state, 
            setMap: elem => setState(createMap(elem)) 
            }}
        >
            <UseOlMap />
        </OLMap.Provider>
    );
}

function UseOlMap() {
    const { map, setMap } = useContext(OLMap);
    const mapRef = useRef(null);
    useEffect(() => {
        setMap(mapRef.current);
    }, []);
    return (
    <>
      <div ref={mapRef} />
      {map && <Layer points={[new Point(routeCoords[0])]} />}
    </>
    );
}

function Layer({ points }) {
    const styles = {
        route: new Style({
            stroke: new Stroke({
                width: 6, color: [237, 212, 0, 0.8]
            }),
        }),
        icon: new Style({
            image: new Icon({
                anchor: [0.5, 1],
                src: 'data/icon.png',
            }),
        }),
        geoMarker: new Style({
            image: new CircleStyle({
                radius: 7,
                fill: new Fill({ color: 'black' }),
                stroke: new Stroke({
                    color: 'white', width: 2,
                }),
            }),
        }),
    };
    const { map } = useContext(OLMap);
    const [layer] = useState(() => {
        const featureLayer = new VectorLayer({
            source: new VectorSource({
                features: points.map((point) => {
                    const feature = new Feature(point);
                    feature.setId(Math.random());
                    console.log(feature);
                    return feature;
                }),
            }),
            style: () => styles.geoMarker,
        });
        map.addLayer(featureLayer);
        return featureLayer;
    });
    useEffect(() => {
        console.log(layer);
    }, [points]);
    return null;
}
