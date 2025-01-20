# MyCliAI Project

## Project Overview
MyCliAI is a Terminal User Interface (TUI) based chat application that provides an enhanced interaction experience with AI models. It focuses on proper text formatting, conversation management, and local session persistence.

## Core Goals
1. Provide a clean and efficient TUI for AI interaction
2. Support markdown formatting and code highlighting
3. Maintain conversation history across sessions
4. Handle multiple AI providers starting with Gemini
5. Provide a seamless command system for enhanced interaction

## Key Design Decisions

### 1. Architecture
- Implement as a TUI application using Bubble Tea
- Modular design with clear separation of concerns
- SQLite for local storage
- Stateless communication with AI providers

### 2. UI Components
- Clear visual distinction between user and AI messages
- Proper text wrapping with UTF-8 support
- Interactive input handling with command support
- Loading state indication
- Clean initialization and shutdown

### 3. Data Management
- Local SQLite database for persistence
- Session-based conversation tracking
- Message history with timestamps
- Command history support

## Implementation Milestones

### Phase 1: Basic Functionality (v0.1.0) ✓
- [x] Basic TUI implementation
- [x] AI provider integration (Gemini)
- [x] Clean text display with proper wrapping
- [x] Input handling and command parsing
- [x] Initial SQLite schema design

### Phase 2: Core Features (v0.2.0)
- [ ] Implement conversation persistence
- [ ] Add session management
- [ ] Improve markdown rendering
- [ ] Add error handling for API failures
- [ ] Implement basic commands (/help, /clear)

### Phase 3: Enhanced Features (v0.3.0)
- [ ] Add configuration file support
- [ ] Implement conversation export
- [ ] Add theme customization
- [ ] Add conversation search
- [ ] Implement chat history browser

### Phase 4: Advanced Features (v0.4.0)
- [ ] Support for different AI models
- [ ] Add streaming responses
- [ ] Implement code execution
- [ ] Add conversation templates
- [ ] Add file upload support

## Technical Specifications

### Core Components
1. **UI System**
   - Bubble Tea based TUI
   - Lipgloss styling
   - UTF-8 support
   - Proper text wrapping

2. **Data Layer**
   - SQLite storage
   - Session management
   - Message persistence

3. **AI Integration**
   - Provider abstraction
   - Context management
   - Error handling

## Dependencies
- Bubble Tea for TUI
- Lipgloss for styling
- SQLite for storage
- Gemini AI SDK

## Current Status
1. Completed basic TUI implementation
2. Established modular architecture
3. Implemented proper text handling
4. Added initial AI integration
5. Set up basic database schema

## Next Steps
1. Implement conversation persistence
2. Add session management
3. Improve markdown support
4. Add command system
5. Enhance error handling

## Notes
- Focus on user experience and clean interface
- Maintain modular design for extensibility
- Emphasize proper text handling and display
- Plan for multiple AI provider support
- Consider international character support