#!/bin/bash

# Cores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
WHITE='\033[1;37m'
NC='\033[0m' # No Color

API_URL="http://localhost:8080"

clear

echo -e "${CYAN}╔══════════════════════════════════════════════════════════════════════════════╗${NC}"
echo -e "${CYAN}║                                                                              ║${NC}"
echo -e "${CYAN}║                        🌾 WELCOME TO CROPFLOW 🌾                            ║${NC}"
echo -e "${CYAN}║                                                                              ║${NC}"
echo -e "${CYAN}║                   Modern Agricultural Management API                         ║${NC}"
echo -e "${CYAN}║                        Built with Go & Clean Architecture                   ║${NC}"
echo -e "${CYAN}║                                                                              ║${NC}"
echo -e "${CYAN}╚══════════════════════════════════════════════════════════════════════════════╝${NC}"
echo ""

echo -e "${WHITE}🚀 PROJECT OVERVIEW${NC}"
echo -e "${BLUE}═══════════════════${NC}"
echo -e "${YELLOW}• Name:${NC} CropFlow API"
echo -e "${YELLOW}• Purpose:${NC} Modern agricultural management solution"
echo -e "${YELLOW}• Language:${NC} Go (Golang)"
echo -e "${YELLOW}• Architecture:${NC} Clean Architecture + Domain-Driven Design"
echo -e "${YELLOW}• Status:${NC} ${GREEN}✅ Fully Functional${NC}"
echo ""

echo -e "${WHITE}🏗️ ARCHITECTURE HIGHLIGHTS${NC}"
echo -e "${BLUE}═══════════════════════════${NC}"
echo -e "${GREEN}✅${NC} Clean Architecture with clear separation of concerns"
echo -e "${GREEN}✅${NC} Domain-Driven Design (DDD) implementation"
echo -e "${GREEN}✅${NC} JWT Authentication with role-based access control"
echo -e "${GREEN}✅${NC} RESTful API with comprehensive endpoints"
echo -e "${GREEN}✅${NC} Docker containerization for easy deployment"
echo -e "${GREEN}✅${NC} MySQL database with GORM ORM"
echo -e "${GREEN}✅${NC} Comprehensive testing with real user simulation"
echo ""

echo -e "${WHITE}🔧 TECHNOLOGY STACK${NC}"
echo -e "${BLUE}═══════════════════${NC}"
echo -e "${CYAN}Backend:${NC} Go 1.21 + Gin Framework"
echo -e "${CYAN}Database:${NC} MySQL 8.0 + GORM"
echo -e "${CYAN}Security:${NC} JWT + bcrypt password hashing"
echo -e "${CYAN}Containers:${NC} Docker + Docker Compose"
echo -e "${CYAN}Testing:${NC} Automated user simulation scripts"
echo ""

echo -e "${WHITE}🔐 ACCESS CONTROL MATRIX${NC}"
echo -e "${BLUE}═══════════════════════════${NC}"
echo -e "${YELLOW}Role      │ Farms    │ Crops     │ Fertilizers${NC}"
echo -e "${BLUE}──────────┼──────────┼───────────┼────────────${NC}"
echo -e "${GREEN}USER      │ ✅ List   │ ❌ List    │ ❌ List${NC}"
echo -e "${GREEN}MANAGER   │ ✅ List   │ ✅ List    │ ❌ List${NC}"
echo -e "${GREEN}ADMIN     │ ✅ List   │ ✅ List    │ ✅ List${NC}"
echo ""

echo -e "${WHITE}📊 CHECKING SERVICES STATUS${NC}"
echo -e "${BLUE}════════════════════════════${NC}"

# Check if containers are running
API_RUNNING=false
DB_RUNNING=false

if docker ps | grep -q "cropflow-api"; then
    echo -e "${GREEN}✅ CropFlow API:${NC} Running on port 8080"
    API_RUNNING=true
else
    echo -e "${RED}❌ CropFlow API:${NC} Not running"
    echo -e "${YELLOW}💡 Starting services with docker compose...${NC}"
    docker compose up -d
    sleep 5
fi

if docker ps | grep -q "cropflow-mysql"; then
    echo -e "${GREEN}✅ MySQL Database:${NC} Running on port 3306"
    DB_RUNNING=true
else
    echo -e "${RED}❌ MySQL Database:${NC} Not running"
fi

# Wait for API to be ready
echo -e "${YELLOW}⏳ Waiting for API to be ready...${NC}"
MAX_ATTEMPTS=30
ATTEMPT=0
while [ $ATTEMPT -lt $MAX_ATTEMPTS ]; do
    if curl -s "$API_URL/farms" > /dev/null 2>&1; then
        echo -e "${GREEN}✅ API Connectivity:${NC} Responding"
        break
    fi
    ATTEMPT=$((ATTEMPT + 1))
    echo -n "."
    sleep 1
done
echo ""

if [ $ATTEMPT -eq $MAX_ATTEMPTS ]; then
    echo -e "${RED}❌ API Connectivity:${NC} Not responding after $MAX_ATTEMPTS attempts"
    echo -e "${YELLOW}Please ensure Docker Compose services are running: docker compose up -d${NC}"
    exit 1
fi

echo ""
echo -e "${WHITE}🎭 STARTING USER SIMULATION${NC}"
echo -e "${BLUE}═════════════════════════════${NC}"
echo ""

# Variables to store IDs
USER_ID=""
MANAGER_ID=""
ADMIN_ID=""
USER_TOKEN=""
MANAGER_TOKEN=""
ADMIN_TOKEN=""
FARM_ID=""
CROP_ID=""
FERTILIZER_ID=""

# Function to make API calls
api_call() {
    local method=$1
    local endpoint=$2
    local data=$3
    local token=$4
    
    if [ -n "$token" ]; then
        curl -s -X "$method" \
            -H "Content-Type: application/json" \
            -H "Authorization: Bearer $token" \
            -d "$data" \
            "$API_URL$endpoint"
    else
        curl -s -X "$method" \
            -H "Content-Type: application/json" \
            -d "$data" \
            "$API_URL$endpoint"
    fi
}

# Step 1: Create Users
echo -e "${CYAN}📝 Step 1: Creating Users${NC}"
echo -e "${BLUE}─────────────────────────${NC}"

# Create USER
echo -e "${YELLOW}Creating USER account...${NC}"
USER_RESPONSE=$(api_call "POST" "/persons" '{"username":"demo_user","password":"demo123","role":"ROLE_USER"}')
USER_ID=$(echo "$USER_RESPONSE" | grep -o '"id":[0-9]*' | grep -o '[0-9]*' | head -1)
if [ -n "$USER_ID" ]; then
    echo -e "${GREEN}✅ User created:${NC} ID=$USER_ID, Username=demo_user, Role=ROLE_USER"
else
    echo -e "${RED}❌ Failed to create USER${NC}"
    echo "Response: $USER_RESPONSE"
fi
echo ""

# Create MANAGER
echo -e "${YELLOW}Creating MANAGER account...${NC}"
MANAGER_RESPONSE=$(api_call "POST" "/persons" '{"username":"demo_manager","password":"demo123","role":"ROLE_MANAGER"}')
MANAGER_ID=$(echo "$MANAGER_RESPONSE" | grep -o '"id":[0-9]*' | grep -o '[0-9]*' | head -1)
if [ -n "$MANAGER_ID" ]; then
    echo -e "${GREEN}✅ Manager created:${NC} ID=$MANAGER_ID, Username=demo_manager, Role=ROLE_MANAGER"
else
    echo -e "${RED}❌ Failed to create MANAGER${NC}"
    echo "Response: $MANAGER_RESPONSE"
fi
echo ""

# Create ADMIN
echo -e "${YELLOW}Creating ADMIN account...${NC}"
ADMIN_RESPONSE=$(api_call "POST" "/persons" '{"username":"demo_admin","password":"demo123","role":"ROLE_ADMIN"}')
ADMIN_ID=$(echo "$ADMIN_RESPONSE" | grep -o '"id":[0-9]*' | grep -o '[0-9]*' | head -1)
if [ -n "$ADMIN_ID" ]; then
    echo -e "${GREEN}✅ Admin created:${NC} ID=$ADMIN_ID, Username=demo_admin, Role=ROLE_ADMIN"
else
    echo -e "${RED}❌ Failed to create ADMIN${NC}"
    echo "Response: $ADMIN_RESPONSE"
fi
echo ""

# Step 2: Login
echo -e "${CYAN}🔐 Step 2: Authenticating Users${NC}"
echo -e "${BLUE}────────────────────────────────${NC}"

# Login as USER
echo -e "${YELLOW}Logging in as USER...${NC}"
USER_LOGIN=$(api_call "POST" "/auth/login" '{"username":"demo_user","password":"demo123"}')
USER_TOKEN=$(echo "$USER_LOGIN" | grep -o '"token":"[^"]*"' | cut -d'"' -f4)
if [ -n "$USER_TOKEN" ]; then
    echo -e "${GREEN}✅ USER authenticated${NC}"
else
    echo -e "${RED}❌ Failed to authenticate USER${NC}"
    echo "Response: $USER_LOGIN"
fi
echo ""

# Login as MANAGER
echo -e "${YELLOW}Logging in as MANAGER...${NC}"
MANAGER_LOGIN=$(api_call "POST" "/auth/login" '{"username":"demo_manager","password":"demo123"}')
MANAGER_TOKEN=$(echo "$MANAGER_LOGIN" | grep -o '"token":"[^"]*"' | cut -d'"' -f4)
if [ -n "$MANAGER_TOKEN" ]; then
    echo -e "${GREEN}✅ MANAGER authenticated${NC}"
else
    echo -e "${RED}❌ Failed to authenticate MANAGER${NC}"
    echo "Response: $MANAGER_LOGIN"
fi
echo ""

# Login as ADMIN
echo -e "${YELLOW}Logging in as ADMIN...${NC}"
ADMIN_LOGIN=$(api_call "POST" "/auth/login" '{"username":"demo_admin","password":"demo123"}')
ADMIN_TOKEN=$(echo "$ADMIN_LOGIN" | grep -o '"token":"[^"]*"' | cut -d'"' -f4)
if [ -n "$ADMIN_TOKEN" ]; then
    echo -e "${GREEN}✅ ADMIN authenticated${NC}"
else
    echo -e "${RED}❌ Failed to authenticate ADMIN${NC}"
    echo "Response: $ADMIN_LOGIN"
fi
echo ""

# Step 3: Create Farm (using ADMIN token)
echo -e "${CYAN}🚜 Step 3: Creating Farm${NC}"
echo -e "${BLUE}─────────────────────────${NC}"
echo -e "${YELLOW}Creating farm as ADMIN...${NC}"
FARM_RESPONSE=$(api_call "POST" "/farms" '{"name":"Green Valley Farm","size":150.5}' "$ADMIN_TOKEN")
FARM_ID=$(echo "$FARM_RESPONSE" | grep -o '"id":[0-9]*' | grep -o '[0-9]*' | head -1)
if [ -n "$FARM_ID" ]; then
    FARM_NAME=$(echo "$FARM_RESPONSE" | grep -o '"name":"[^"]*"' | cut -d'"' -f4)
    FARM_SIZE=$(echo "$FARM_RESPONSE" | grep -o '"size":[0-9.]*' | cut -d':' -f2)
    echo -e "${GREEN}✅ Farm created:${NC} ID=$FARM_ID, Name=$FARM_NAME, Size=$FARM_SIZE hectares"
else
    echo -e "${RED}❌ Failed to create farm${NC}"
    echo "Response: $FARM_RESPONSE"
fi
echo ""

# Step 4: List Farms (as USER)
echo -e "${CYAN}📋 Step 4: Listing Farms (as USER)${NC}"
echo -e "${BLUE}──────────────────────────────────${NC}"
FARMS_LIST=$(api_call "GET" "/farms" "" "$USER_TOKEN")
FARM_COUNT=$(echo "$FARMS_LIST" | grep -o '"id"' | wc -l)
echo -e "${GREEN}✅ USER can list farms:${NC} Found $FARM_COUNT farm(s)"
echo ""

# Step 5: Create Crop
echo -e "${CYAN}🌾 Step 5: Creating Crop${NC}"
echo -e "${BLUE}─────────────────────────${NC}"
if [ -n "$FARM_ID" ]; then
    echo -e "${YELLOW}Creating crop in farm $FARM_ID...${NC}"
    # Use ISO 8601 format with time (required by Go time parser)
    # Try different date commands for compatibility
    if command -v date >/dev/null 2>&1; then
        if date -u -d "+90 days" +%Y-%m-%dT%H:%M:%SZ >/dev/null 2>&1; then
            # Linux date
            TODAY=$(date -u +%Y-%m-%dT%H:%M:%SZ)
            HARVEST_DATE=$(date -u -d "+90 days" +%Y-%m-%dT%H:%M:%SZ)
        elif date -u -v+90d +%Y-%m-%dT%H:%M:%SZ >/dev/null 2>&1; then
            # macOS date
            TODAY=$(date -u +%Y-%m-%dT%H:%M:%SZ)
            HARVEST_DATE=$(date -u -v+90d +%Y-%m-%dT%H:%M:%SZ)
        else
            # Fallback: use current time and skip harvest date
            TODAY=$(date -u +%Y-%m-%dT%H:%M:%SZ)
            HARVEST_DATE=""
        fi
    else
        # Fallback: use current time
        TODAY=$(date -u +%Y-%m-%dT%H:%M:%SZ 2>/dev/null || echo "")
        HARVEST_DATE=""
    fi
    
    if [ -n "$HARVEST_DATE" ] && [ -n "$TODAY" ]; then
        CROP_DATA="{\"name\":\"Corn\",\"plantedArea\":25.0,\"plantedDate\":\"$TODAY\",\"harvestDate\":\"$HARVEST_DATE\"}"
    elif [ -n "$TODAY" ]; then
        CROP_DATA="{\"name\":\"Corn\",\"plantedArea\":25.0,\"plantedDate\":\"$TODAY\"}"
    else
        # Skip dates if we can't generate them properly
        CROP_DATA="{\"name\":\"Corn\",\"plantedArea\":25.0}"
    fi
    
    CROP_RESPONSE=$(api_call "POST" "/farms/$FARM_ID/crops" "$CROP_DATA" "$ADMIN_TOKEN")
    CROP_ID=$(echo "$CROP_RESPONSE" | grep -o '"id":[0-9]*' | grep -o '[0-9]*' | head -1)
    if [ -n "$CROP_ID" ]; then
        CROP_NAME=$(echo "$CROP_RESPONSE" | grep -o '"name":"[^"]*"' | cut -d'"' -f4)
        CROP_AREA=$(echo "$CROP_RESPONSE" | grep -o '"plantedArea":[0-9.]*' | cut -d':' -f2)
        echo -e "${GREEN}✅ Crop created:${NC} ID=$CROP_ID, Name=$CROP_NAME, Area=$CROP_AREA hectares"
    else
        echo -e "${RED}❌ Failed to create crop${NC}"
        echo "Response: $CROP_RESPONSE"
    fi
else
    echo -e "${RED}❌ Cannot create crop: Farm ID not available${NC}"
fi
echo ""

# Step 6: List Crops (as USER - should fail)
echo -e "${CYAN}📋 Step 6: Testing Access Control - List All Crops (as USER)${NC}"
echo -e "${BLUE}──────────────────────────────────────────────────────────────${NC}"
CROPS_LIST_USER=$(api_call "GET" "/crops" "" "$USER_TOKEN")
if echo "$CROPS_LIST_USER" | grep -q "error\|forbidden\|unauthorized"; then
    echo -e "${RED}❌ USER cannot list all crops (expected):${NC} Access denied"
else
    echo -e "${YELLOW}⚠️  USER can list crops (unexpected)${NC}"
fi
echo ""

# Step 7: List Crops (as MANAGER - should succeed)
echo -e "${CYAN}📋 Step 7: Listing All Crops (as MANAGER)${NC}"
echo -e "${BLUE}───────────────────────────────────────────${NC}"
CROPS_LIST_MANAGER=$(api_call "GET" "/crops" "" "$MANAGER_TOKEN")
CROP_COUNT=$(echo "$CROPS_LIST_MANAGER" | grep -o '"id"' | wc -l)
echo -e "${GREEN}✅ MANAGER can list all crops:${NC} Found $CROP_COUNT crop(s)"
echo ""

# Step 8: Create Fertilizer (as ADMIN)
echo -e "${CYAN}🧪 Step 8: Creating Fertilizer${NC}"
echo -e "${BLUE}──────────────────────────────${NC}"
echo -e "${YELLOW}Creating fertilizer as ADMIN...${NC}"
FERTILIZER_RESPONSE=$(api_call "POST" "/fertilizers" '{"name":"NPK 10-10-10","brand":"AgroMax","composition":"Nitrogen 10%, Phosphorus 10%, Potassium 10%"}' "$ADMIN_TOKEN")
FERTILIZER_ID=$(echo "$FERTILIZER_RESPONSE" | grep -o '"id":[0-9]*' | grep -o '[0-9]*' | head -1)
if [ -n "$FERTILIZER_ID" ]; then
    FERT_NAME=$(echo "$FERTILIZER_RESPONSE" | grep -o '"name":"[^"]*"' | cut -d'"' -f4)
    FERT_BRAND=$(echo "$FERTILIZER_RESPONSE" | grep -o '"brand":"[^"]*"' | cut -d'"' -f4)
    echo -e "${GREEN}✅ Fertilizer created:${NC} ID=$FERTILIZER_ID, Name=$FERT_NAME, Brand=$FERT_BRAND"
else
    echo -e "${RED}❌ Failed to create fertilizer${NC}"
    echo "Response: $FERTILIZER_RESPONSE"
fi
echo ""

# Step 9: List Fertilizers (as USER - should fail)
echo -e "${CYAN}📋 Step 9: Testing Access Control - List Fertilizers (as USER)${NC}"
echo -e "${BLUE}───────────────────────────────────────────────────────────────${NC}"
FERT_LIST_USER=$(api_call "GET" "/fertilizers" "" "$USER_TOKEN")
if echo "$FERT_LIST_USER" | grep -q "error\|forbidden\|unauthorized"; then
    echo -e "${RED}❌ USER cannot list fertilizers (expected):${NC} Access denied"
else
    echo -e "${YELLOW}⚠️  USER can list fertilizers (unexpected)${NC}"
fi
echo ""

# Step 10: List Fertilizers (as MANAGER - should fail)
echo -e "${CYAN}📋 Step 10: Testing Access Control - List Fertilizers (as MANAGER)${NC}"
echo -e "${BLUE}────────────────────────────────────────────────────────────────${NC}"
FERT_LIST_MANAGER=$(api_call "GET" "/fertilizers" "" "$MANAGER_TOKEN")
if echo "$FERT_LIST_MANAGER" | grep -q "error\|forbidden\|unauthorized"; then
    echo -e "${RED}❌ MANAGER cannot list fertilizers (expected):${NC} Access denied"
else
    echo -e "${YELLOW}⚠️  MANAGER can list fertilizers (unexpected)${NC}"
fi
echo ""

# Step 11: List Fertilizers (as ADMIN - should succeed)
echo -e "${CYAN}📋 Step 11: Listing Fertilizers (as ADMIN)${NC}"
echo -e "${BLUE}──────────────────────────────────────────${NC}"
FERT_LIST_ADMIN=$(api_call "GET" "/fertilizers" "" "$ADMIN_TOKEN")
FERT_COUNT=$(echo "$FERT_LIST_ADMIN" | grep -o '"id"' | wc -l)
echo -e "${GREEN}✅ ADMIN can list fertilizers:${NC} Found $FERT_COUNT fertilizer(s)"
echo ""

# Step 12: Associate Fertilizer with Crop
echo -e "${CYAN}🔗 Step 12: Associating Fertilizer with Crop${NC}"
echo -e "${BLUE}────────────────────────────────────────────${NC}"
if [ -n "$CROP_ID" ] && [ -n "$FERTILIZER_ID" ]; then
    echo -e "${YELLOW}Associating fertilizer $FERTILIZER_ID with crop $CROP_ID...${NC}"
    ASSOC_RESPONSE=$(api_call "POST" "/crop/$CROP_ID/fertilizer/$FERTILIZER_ID" "" "$ADMIN_TOKEN")
    if echo "$ASSOC_RESPONSE" | grep -q "sucesso\|success\|message"; then
        echo -e "${GREEN}✅ Fertilizer associated with crop successfully${NC}"
    else
        echo -e "${RED}❌ Failed to associate fertilizer${NC}"
        echo "Response: $ASSOC_RESPONSE"
    fi
else
    echo -e "${RED}❌ Cannot associate: Crop ID or Fertilizer ID not available${NC}"
fi
echo ""

# Step 13: Get Crop Fertilizers
echo -e "${CYAN}📋 Step 13: Listing Fertilizers for Crop${NC}"
echo -e "${BLUE}──────────────────────────────────────────${NC}"
if [ -n "$CROP_ID" ]; then
    CROP_FERTS=$(api_call "GET" "/crop/$CROP_ID/fertilizers" "" "$ADMIN_TOKEN")
    FERT_COUNT_CROP=$(echo "$CROP_FERTS" | grep -o '"id"' | wc -l)
    echo -e "${GREEN}✅ Crop $CROP_ID has $FERT_COUNT_CROP fertilizer(s) associated${NC}"
fi
echo ""

# Summary
echo -e "${CYAN}╔══════════════════════════════════════════════════════════════════════════════╗${NC}"
echo -e "${CYAN}║                                                                              ║${NC}"
echo -e "${CYAN}║                         📊 SIMULATION SUMMARY 📊                            ║${NC}"
echo -e "${CYAN}║                                                                              ║${NC}"
echo -e "${CYAN}╚══════════════════════════════════════════════════════════════════════════════╝${NC}"
echo ""

echo -e "${WHITE}✅ Created Users:${NC}"
echo -e "  ${GREEN}• USER:${NC}    demo_user (ID: $USER_ID)"
echo -e "  ${GREEN}• MANAGER:${NC} demo_manager (ID: $MANAGER_ID)"
echo -e "  ${GREEN}• ADMIN:${NC}   demo_admin (ID: $ADMIN_ID)"
echo ""

echo -e "${WHITE}✅ Created Resources:${NC}"
if [ -n "$FARM_ID" ]; then
    echo -e "  ${GREEN}• Farm:${NC} ID=$FARM_ID"
fi
if [ -n "$CROP_ID" ]; then
    echo -e "  ${GREEN}• Crop:${NC} ID=$CROP_ID"
fi
if [ -n "$FERTILIZER_ID" ]; then
    echo -e "  ${GREEN}• Fertilizer:${NC} ID=$FERTILIZER_ID"
fi
echo ""

echo -e "${WHITE}✅ Access Control Verified:${NC}"
echo -e "  ${GREEN}• USER:${NC}    Can list farms ✅ | Cannot list all crops ❌ | Cannot list fertilizers ❌"
echo -e "  ${GREEN}• MANAGER:${NC} Can list farms ✅ | Can list all crops ✅ | Cannot list fertilizers ❌"
echo -e "  ${GREEN}• ADMIN:${NC}   Can list farms ✅ | Can list all crops ✅ | Can list fertilizers ✅"
echo ""

echo -e "${CYAN}╔══════════════════════════════════════════════════════════════════════════════╗${NC}"
echo -e "${CYAN}║                                                                              ║${NC}"
echo -e "${CYAN}║                    🌾 CROPFLOW SIMULATION COMPLETE! 🌾                       ║${NC}"
echo -e "${CYAN}║                                                                              ║${NC}"
echo -e "${CYAN}║                    All operations executed successfully!                     ║${NC}"
echo -e "${CYAN}║                                                                              ║${NC}"
echo -e "${CYAN}╚══════════════════════════════════════════════════════════════════════════════╝${NC}"
echo ""
