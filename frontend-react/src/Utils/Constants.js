const DEFAULT_COLOR = {
   red: '#FFCACA',
   yellow: '#FDFFB1',
   green: '#CDFFCD'
};

export const BEGIN_STATE = {
   umlImgName: 'begin',
   color: DEFAULT_COLOR
};


export const END_STATE = {
   umlImgName: 'end',
   color: DEFAULT_COLOR
};

export const WORKING_STATE = {
   red: {
      umlImgName: 'red',
      color: {
         red: '#cc3232',
         yellow: '#FDFFB1',
         green: '#CDFFCD'
      }
   },
   green: {
      umlImgName: 'green',
      color: {
         red: '#FFCACA',
         yellow: '#FDFFB1',
         green: '#2dc937'
      }
   },
   yellow: {
      umlImgName: 'yellow',
      color: {
         red: '#FFCACA',
         yellow: '#e7b416',
         green: '#CDFFCD'
      }
   }
};

export const SYSTEM_ERROR_STATE = {
   default: {
      umlImgName: 'flashingRed',
      color: DEFAULT_COLOR
   },
   red: {
      umlImgName: 'flashingRed',
      color: {
         red: '#cc3232',
         yellow: '#FDFFB1',
         green: '#CDFFCD'
      }
   }
};


export const STATES = {
   'INITAL_STATE': 'initial',
   'BEGIN_STATE': 'begin',
   'WORKING_STATE': 'working',
   'ERROR_STATE': 'error',
   'END_STATE': 'end'
};